<template>
  <div class="comment-wrapper">
    <header>
      <h2>{{ comments.length }}件のコメント</h2>
      <button @click="close">
        <font-awesome-icon icon="times" />
      </button>
    </header>
    <div class="comment">
      <div v-show="showInput" class="comment-input">
        <BaseAvatar :src="avatar" :size="50" />
        <div class="comment-input-left">
          <textarea
            ref="commentInput"
            placeholder="コメントを入力..."
          ></textarea>
          <div class="comment-input-action">
            <button class="cancel-btn" @click="clearTextArea">
              キャンセル
            </button>
            <button class="comment-btn" @click="comment">コメント</button>
          </div>
        </div>
      </div>
      <div v-for="c in comments" :key="c.comment_id" class="comment-item">
        <BaseAvatar :src="c.author.avatar_url" :size="40" />
        <div class="comment-item-right">
          <div>{{ c.author.username }}</div>
          <div>
            {{ c.body }}
          </div>
          <div>{{ c.created_at }}</div>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { $commentApi } from '@/plugins/api'
import { Comment } from '@/api-client-axios'
import BaseAvatar from '@/components/BaseAvatar.vue'
import { meStore } from '@/store'

interface Data {
  articleId: string
  comments: Array<Comment>
  error: boolean
}

export default Vue.extend({
  components: {
    BaseAvatar,
  },
  async asyncData({ params }): Promise<Data | void> {
    try {
      // コメント取得
      const { data } = await $commentApi().getComments(params.id)
      console.log(data)
      return {
        articleId: params.id,
        comments: data,
        error: false,
      }
    } catch {
      return {
        articleId: params.id,
        comments: [],
        error: true,
      }
    }
  },
  data(): Data {
    return {} as Data
  },
  computed: {
    showInput() {
      return meStore.getMe !== null
    },
    avatar() {
      return meStore.getMe?.avatar_url || '/default.png'
    },
  },
  methods: {
    close() {
      window.history.back()
    },
    clearTextArea() {
      const textarea = this.$refs.commentInput as HTMLTextAreaElement
      textarea.value = ''
    },
    async comment() {
      const val = (this.$refs.commentInput as HTMLTextAreaElement).value
      const { data } = await $commentApi().postComment(this.articleId, {
        body: val,
      })
      const tmp: Comment[] = [data]
      this.comments = tmp.concat(this.comments)
      this.clearTextArea()
    },
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';

.comment-wrapper {
  background-color: white;
  position: absolute;
  width: 100vw;
  top: 0;
  right: 0;
  .comment {
    max-width: 520px;
    margin: 0 auto;
  }
  header {
    display: flex;
    height: 40px;
    padding: 0 10px;
    background-color: $bg-blue;
    align-items: center;
    h2 {
      display: block;
      color: white;
    }
    button {
      display: block;
      margin-left: auto;
      color: white;
    }
  }
  .comment-avatar {
    background-color: grey;
    border-radius: 50%;
    width: 50px;
    height: 50px;
  }
  .comment-input {
    display: flex;
    border-bottom: 1px solid $stroke-grey;
    padding: 10px 5px;
    textarea {
      margin-left: 20px;
      height: 60px;
      padding: 10px 0;
      font-size: 1.1rem;
      resize: none;
      width: 100%;
    }
  }
  .comment-input-left {
    width: calc(100% - 50px - 20px);
  }
  .comment-input-action {
    display: flex;
    .cancel-btn {
      display: block;
      margin-left: auto;
      color: $stroke-grey;
    }
    .comment-btn {
      display: block;
      color: $btn-blue;
      margin-left: 15px;
    }
  }

  .comment-item {
    display: flex;
    border-bottom: 1px solid $stroke-grey;
    padding: 15px 0;
    margin: 10px 30px;
    min-height: 100px;
    .comment-item-right {
      margin-left: 20px;
      width: calc(100% - 50px - 30px);
      display: flex;
      flex-direction: column;
      justify-content: space-between;
      div:nth-child(1) {
        font-weight: bold;
      }
      div:nth-child(2) {
        margin: 10px 0;
      }
      div:nth-child(3) {
        color: $stroke-grey;
        margin-left: auto;
      }
    }
  }
}
</style>
