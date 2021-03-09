<template>
  <div id="new-article">
    <div class="new-article-container">
      <div id="toolbar">
        <button class="ql-header" value="2"></button>
        <button class="ql-image" value="2"></button>
        <button ref="youtubeBtn">
          <font-awesome-icon :icon="['fab', 'youtube']" style="color: #444" />
        </button>
        <button id="post-btn" ref="postBtn" @click="uploadArticle">投稿</button>
        <input
          v-show="false"
          ref="inputImage"
          type="file"
          accept="image/png, image/jpeg"
          @change="uploadImage"
        />
      </div>
      <div class="thumbnail-empty">
        <button v-show="!showThumbnail" @click="selectThumbnail">
          <font-awesome-icon :icon="['fas', 'image']" />
          <div>サムネイルを追加</div>
        </button>
        <input
          v-show="false"
          ref="inputThumbnail"
          type="file"
          accept="image/png, image/jpeg"
          @change="uploadThumbnail"
        />
      </div>
      <div v-show="showThumbnail" class="thumbnail" @click="selectThumbnail">
        <img :src="thumbnail" />
      </div>
      <button @click="titleFocus">title</button>
      <button @click="bodyFocus">body</button>
      <div id="title-editor"></div>
      <div id="body-editor"></div>
    </div>
    <VideoUrlInput
      v-if="showVideoUrlInput"
      :on-click-cancel="closeVideoUrlInput"
      :on-click-ok="embedVideo"
      :on-input="updateVideoUrl"
    />
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { $articleApi, $imageApi } from '@/plugins/api'
import 'quill/dist/quill.snow.css'
import VideoUrlInput from '@/components/VideoUrlInput.vue'
import { baseModalState } from '~/store'

interface data {
  thumbnail: string
  title: string
  body: string
  titleEditor: any
  bodyEditor: any
  showVideoUrlInput: boolean
  videoUrl: string
}

export default Vue.extend({
  components: {
    VideoUrlInput,
  },

  data(): data {
    return {
      thumbnail: '',
      title: '',
      body: '',
      titleEditor: null,
      bodyEditor: null,
      showVideoUrlInput: false,
      videoUrl: '',
    }
  },
  computed: {
    showThumbnail(): boolean {
      return !!this.thumbnail
    },
  },
  async beforeMount() {
    const Quill = (await import('quill')).default
    const BlockEmbed = Quill.import('blots/block/embed')

    // タイトルエディター
    this.titleEditor = new Quill('#title-editor', {
      formats: [],
      placeholder: 'タイトル',
    })
    // 改行が入力されたら削除
    this.titleEditor.on('text-change', (delta: any, _: any, __: any) => {
      if (delta.ops[delta.ops.length - 1].insert === '\n') {
        this.titleEditor.deleteText(this.titleEditor.getLength() - 1, 1)
      }
    })
    // 貼り付け時にplain textにする
    this.titleEditor.clipboard.addMatcher(
      Node.ELEMENT_NODE,
      (node: any, _: any) => {
        const plaintext = node.textContent
        const Delta = Quill.import('delta')
        return new Delta().insert(plaintext)
      }
    )
    class VideoBlot extends BlockEmbed {
      static create(url: string) {
        const node = super.create() as HTMLElement
        const child = node.appendChild(
          document.createElement('iframe')
        ) as HTMLIFrameElement
        node.setAttribute('class', 'iframe-wrapper')
        child.setAttribute('src', url)
        child.setAttribute('frameborder', '0')
        child.setAttribute('allowfullscreen', 'true')
        return node
      }

      static formats(node: HTMLElement) {
        const format = {}
        if (node.hasAttribute('height')) {
          // @ts-ignore
          format.height = node.getAttribute('height')
        }
        if (node.hasAttribute('width')) {
          // @ts-ignore
          format.width = node.getAttribute('width')
        }
        return format
      }

      static value(node: HTMLElement) {
        return node.getAttribute('src')
      }

      format(name: string, value: string) {
        if (name === 'height' || name === 'width') {
          if (value) {
            this.domNode.setAttribute(name, value)
          } else {
            this.domNode.removeAttribute(name, value)
          }
        } else {
          super.format(name, value)
        }
      }
    }
    VideoBlot.blotName = 'video'
    VideoBlot.tagName = 'div'
    Quill.register(VideoBlot)

    // ボディエディター
    this.bodyEditor = new Quill('#body-editor', {
      modules: {
        toolbar: '#toolbar',
      },
      formats: ['link', 'header', 'image', 'video'],
      placeholder: '今日はゲームでどんなことがありましたか？',
      theme: 'snow',
    })
    this.bodyEditor.getModule('toolbar').addHandler('image', this.selectImage)
    // 貼り付け時にplain textにする
    this.bodyEditor.clipboard.addMatcher(
      Node.ELEMENT_NODE,
      (node: any, _: any) => {
        const plaintext = node.textContent
        const Delta = Quill.import('delta')
        return new Delta().insert(plaintext)
      }
    )
    const youtubeBtn = this.$refs.youtubeBtn as HTMLButtonElement
    youtubeBtn.onclick = () => {
      this.showVideoUrlInput = true
    }
  },

  methods: {
    async uploadArticle() {
      try {
        // ボタンをdisabledにする
        const btn = this.$refs.postBtn as HTMLButtonElement
        btn.setAttribute('disabled', 'disabled')

        // 記事投稿APIを叩く
        this.title = this.titleEditor.getText()
        this.body = this.bodyEditor.container.firstChild.innerHTML
        const { data } = await $articleApi().postArticle({
          thumbnail_url: this.thumbnail,
          title: this.title,
          body: this.body,
        })

        // 投稿に成功したら、記事ページに遷移
        if (data && data.article_id) {
          this.$router.push(`/articles/${data.article_id}`)
        }
      } catch (err) {
        baseModalState.setModal({
          showModal: true,
          message: '記事の投稿に失敗しました',
        })
      } finally {
        // ボタンのdisabledを解除
        const btn = this.$refs.postBtn as HTMLButtonElement
        btn.removeAttribute('disabled')
      }
    },

    selectImage() {
      if (document !== undefined && document !== null) {
        const e = this.$refs.inputImage as HTMLInputElement
        if (e !== null) {
          e.click()
        }
      }
    },
    async uploadImage(event: Event) {
      try {
        const files = (event.target as HTMLInputElement).files
        if (files && files.length) {
          const file = files[0]
          if (file.size > 100000000) {
            throw new Error('サイズオーバー')
          }
          const result = await $imageApi().uploadImage(file)
          this.bodyEditor.format('image', result.data.url)
        }
      } catch (err) {}
    },

    selectThumbnail() {
      if (document !== undefined && document !== null) {
        const e = this.$refs.inputThumbnail as HTMLInputElement
        if (e !== null) {
          e.click()
        }
      }
    },
    async uploadThumbnail(event: Event) {
      try {
        const files = (event.target as HTMLInputElement).files
        if (files && files.length) {
          const file = files[0]
          if (file.size > 100000000) {
            throw new Error('サイズオーバー')
          }
          const result = await $imageApi().uploadImage(file)
          this.thumbnail = result.data.url
        }
      } catch (err) {}
    },

    closeVideoUrlInput() {
      this.showVideoUrlInput = false
      this.videoUrl = ''
    },
    embedVideo() {
      try {
        const range = this.bodyEditor.getSelection(true)
        this.bodyEditor.insertText(range.index, '\n', 'user')
        let url = ''
        const urlObj = new URL(this.videoUrl)
        if (urlObj.host === 'www.youtube.com') {
          const params = urlObj.searchParams
          const videoId = params.get('v')
          url = `https://www.youtube.com/embed/${videoId}?showinfo=0`
        }
        this.bodyEditor.insertEmbed(range.index + 1, 'video', url, 'user')
        this.bodyEditor.setSelection(range.index + 2, 0, 'silent')
        this.closeVideoUrlInput()
        const docElm = document.documentElement
        window.scroll(0, docElm.scrollHeight - docElm.clientHeight)
      } catch {
        this.closeVideoUrlInput()
      }
    },
    updateVideoUrl(url: string) {
      this.videoUrl = url
    },
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';
#new-article {
  background-color: white;
  max-width: 640px;
  margin: 0 auto;
  border-radius: 10px 10px 0 0;
}

.new-article-container {
  max-width: 640px;
  margin: 0 auto;
}

#toolbar {
  border: none;
  border-bottom: 1px solid $base-grey;
  height: 45px;
  position: sticky;
  top: 0;
  background-color: white;
  z-index: 1;
  button {
    height: 100%;
  }
  #post-btn {
    color: white;
    font-weight: bold;
    text-align: center;
    padding: 5px 0;
    border-radius: 15px;
    margin: 0 10px;
    width: 60px;
    background-color: $btn-red;
    float: right;
  }
  #post-btn:disabled {
    background-color: $stroke-grey;
    cursor: default;
  }
}
.thumbnail-empty {
  color: $base-grey;
  // margin: 50px 0;
  button {
    display: flex;
    width: 100%;
    height: 150px;
    justify-content: center;
    // margin: 0 auto;
    svg {
      margin: auto 0;
      width: 50px;
      height: 50px;
    }
    div {
      margin: auto 0 auto 15px;
    }
  }
}
.thumbnail {
  width: 100%;
  padding-top: 56.25%;
  overflow: hidden;
  position: relative;
  cursor: pointer;

  img {
    position: absolute;
    top: 50%;
    left: 50%;
    width: 100%;
    transform: translate(-50%, -50%);
  }
}
</style>

<style lang="scss">
@import 'assets/global';

#title-editor {
  border: none;
  margin-top: 15px;
  .ql-editor {
    padding-top: 0;
    padding-bottom: 0;
  }
  p {
    font-size: $fs-title;
    font-weight: bold;
  }
  .ql-editor.ql-blank::before {
    font-size: $fs-title;
    font-weight: bold;
    color: $base-grey;
    font-style: normal;
  }
}
#body-editor {
  height: auto;
  border: none;
  padding-bottom: 50vh;
  .ql-editor {
    padding-top: 0;
    padding-bottom: 0;
  }
  .ql-editor.ql-blank::before {
    color: $base-grey;
    font-style: normal;
    margin-top: 10px;
  }
  .ql-editor p {
    margin: 10px 0;
    font-size: $fs-normal;
  }
  .ql-editor h2 {
    margin: 10px 0;
    font-size: $fs-middle;
    font-weight: bold;
  }
  .ql-editor img {
    width: 80%;
    display: block;
    margin: 0 auto;
  }
  .ql-editor .iframe-wrapper {
    position: relative;
    padding-bottom: 56.25%;
    height: 0;
    overflow: hidden;
  }
  .ql-editor iframe {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
  }
}
</style>
