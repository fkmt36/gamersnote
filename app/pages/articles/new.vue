<template>
  <div id="new-article">
    <div class="new-article-container">
      <div id="toolbar">
        <button class="ql-header" value="2"></button>
        <button class="ql-image" value="2"></button>
        <button id="post-btn" @click="uploadArticle">投稿</button>
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
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { $articleApi, $imageApi } from '@/plugins/api'
import 'quill/dist/quill.snow.css'

interface data {
  thumbnail: string
  title: string
  body: string
  titleEditor: any
  bodyEditor: any
}

export default Vue.extend({
  data(): data {
    return {
      thumbnail: '',
      title: '',
      body: '',
      titleEditor: null,
      bodyEditor: null,
    }
  },
  computed: {
    showThumbnail(): boolean {
      return !!this.thumbnail
    },
  },
  async created() {
    try {
      const Quill = (await import('quill')).default
      this.bodyEditor = new Quill('#body-editor', {
        modules: {
          toolbar: '#toolbar',
        },
        formats: ['link', 'header', 'image'],
        placeholder: '今日はゲームでどんなことがありましたか？',
        theme: 'snow',
      })
      this.bodyEditor.getModule('toolbar').addHandler('image', this.selectImage)

      this.titleEditor = new Quill('#title-editor', {
        formats: [],
        placeholder: 'タイトル',
      })
      this.titleEditor.on('text-change', (delta: any, _: any, __: any) => {
        if (delta.ops[delta.ops.length - 1].insert === '\n') {
          this.titleEditor.deleteText(this.titleEditor.getLength() - 1, 1)
        }
      })
    } catch (e) {
      console.error(e)
    }
  },
  methods: {
    async uploadArticle() {
      try {
        this.title = this.titleEditor.getText()
        this.body = this.bodyEditor.container.firstChild.innerHTML
        const { data } = await $articleApi().postArticle({
          thumbnail_url: this.thumbnail,
          title: this.title,
          body: this.body,
        })
        if (data && data.article_id) {
          this.$router.push(`/articles/${data.article_id}`)
        }
      } catch (err) {
      } finally {
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
    // margin-left: auto;
    float: right;
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
  height: 100vh;
  border: none;
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
}
</style>
