<template>
  <div id="edit-article">
    <div class="edit-article-container">
      <div id="toolbar">
        <button class="ql-header" value="2"></button>
        <button class="ql-image" value="2"></button>
        <button ref="youtubeBtn">
          <font-awesome-icon :icon="['fab', 'youtube']" style="color: #444" />
        </button>
        <button id="save-btn" ref="postBtn" @click="saveArticle">保存</button>
        <button id="delete-btn" ref="postBtn" @click="deleteArticle">
          削除
        </button>
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
      <div v-show="showThumbnail" class="thumbnail">
        <img :src="thumbnail" />
      </div>
      <div id="title-editor"></div>
      <div id="body-editor"></div>
    </div>
    <VideoUrlInput
      v-show="showVideoUrlInput"
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
import { baseModalState, meStore } from '~/store'

interface Data {
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

  async asyncData({ route, redirect }): Promise<Data | void> {
    const result = await $articleApi().getArticle(route.params.id as string)
    const me = meStore.getMe
    if (me === null) {
      return redirect('/signin?from=' + route.path)
    }
    if (me.user_id !== result.data.author?.user_id) {
      return redirect('/articles' + result.data.article_id)
    }
    return {
      thumbnail: result.data.thumbnail_url,
      title: result.data.title,
      body: result.data.body,
      titleEditor: null,
      bodyEditor: null,
      showVideoUrlInput: false,
      videoUrl: '',
    }
  },

  data(): Data {
    return {} as Data
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
    // タイトル挿入
    // this.titleEditor.container.innerHTML = '<p>' + this.title + '</p>'
    this.titleEditor.clipboard.dangerouslyPasteHTML(0, this.title)
    // 改行が入力されたら削除
    this.titleEditor.on('text-change', (delta: any, _: any, __: any) => {
      if (delta.ops[delta.ops.length - 1].insert === '\n') {
        this.titleEditor.deleteText(this.titleEditor.getLength() - 1, 1)
      }
    })

    class VideoBlot extends BlockEmbed {
      static create(url: string) {
        const node = super.create(url) as HTMLElement
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
        return node.firstElementChild?.getAttribute('src')
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
    this.bodyEditor.clipboard.dangerouslyPasteHTML(0, this.body)
    const youtubeBtn = this.$refs.youtubeBtn as HTMLButtonElement
    youtubeBtn.onclick = () => {
      this.showVideoUrlInput = true
    }
  },

  methods: {
    async saveArticle() {
      try {
        // ボタンをdisabledにする
        const btn = this.$refs.postBtn as HTMLButtonElement
        btn.setAttribute('disabled', 'disabled')

        // タイトルは1文字以上か
        this.title = this.titleEditor.getText()
        if (this.title.length < 2) {
          console.log('hit')
          throw new Error('Too short title')
        }

        // ボディは1文字以上か
        this.body = this.bodyEditor.container.firstChild.innerHTML
        if (this.bodyEditor.getText().length < 2) {
          throw new Error('Too short body')
        }

        // 記事更新APIを叩く
        const articleId = this.$route.params.id as string
        await $articleApi().putArticle(articleId, {
          thumbnail_url: this.thumbnail,
          title: this.title,
          body: this.body,
        })

        baseModalState.setModal({
          showModal: true,
          message: '記事を更新しました',
        })
      } catch (err) {
        const msg = err.message
        if (msg === 'Too short title') {
          baseModalState.setModal({
            showModal: true,
            message: 'タイトルを入力してください',
          })
        } else if (msg === 'Too short body') {
          baseModalState.setModal({
            showModal: true,
            message: '記事の内容を入力してください',
          })
        } else {
          baseModalState.setModal({
            showModal: true,
            message: '記事の更新に失敗しました',
          })
        }
      } finally {
        // ボタンのdisabledを解除
        const btn = this.$refs.postBtn as HTMLButtonElement
        btn.removeAttribute('disabled')
      }
    },

    async deleteArticle() {
      try {
        // ボタンをdisabledにする
        const btn = this.$refs.postBtn as HTMLButtonElement
        btn.setAttribute('disabled', 'disabled')

        // 記事削除APIを叩く
        const articleId = this.$route.params.id as string
        console.log(articleId)
        await $articleApi().deleteArticle(articleId)

        this.$router.push(`/`)
      } catch (err) {
        console.error(err)
        baseModalState.setModal({
          showModal: true,
          message: '記事の削除に失敗しました',
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
    },
    embedVideo() {
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
    },
    updateVideoUrl(url: string) {
      this.videoUrl = url
    },
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';
#edit-article {
  background-color: white;
  max-width: 640px;
  margin: 0 auto;
  border-radius: 10px 10px 0 0;
}

.edit-article-container {
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

  #save-btn,
  #delete-btn {
    font-weight: bold;
    text-align: center;
    padding: 5px 0;
    border-radius: 15px;
    margin: 0 10px;
    width: 60px;
    float: right;
  }
  #save-btn {
    color: white;
    background-color: $btn-red;
  }
  #delete-btn {
    border: 1px solid $stroke-grey;
    color: $stroke-grey;
    background-color: white;
  }
  #save-btn:disabled,
  #delete-btn:disabled {
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
