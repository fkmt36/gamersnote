<template>
  <div class="input-video-url-wrapper">
    <div class="input-video-url">
      <h2>動画URLを貼り付けてください</h2>
      <p>YouTubeのみ対応しています。</p>
      <input
        type="text"
        placeholder="例: https://www.youtube.com/watch?v=***"
        @input="onInput($event.target.value)"
      />
      <div class="actions">
        <button class="cancel" @click="onClickCancel">キャンセル</button>
        <button class="ok" @click="onClickOk">OK</button>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue, { PropOptions } from 'vue'

export default Vue.extend({
  name: 'VideoUrlInput',

  props: {
    onClickOk: {
      type: Function,
      required: true,
      default: () => {},
    } as PropOptions<() => void>,
    onClickCancel: {
      type: Function,
      required: true,
      default: () => {},
    } as PropOptions<() => void>,
    onInput: {
      type: Function,
      required: true,
      default: () => {},
    } as PropOptions<(v: string) => void>,
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';
.input-video-url-wrapper {
  background-color: rgba(0, 0, 0, 0.4);
  position: fixed;
  height: 100vh;
  width: 100vw;
  top: 0;
  right: 0;
  z-index: 20;
}
.input-video-url {
  padding: 10px 20px;
  max-width: 450px;
  background-color: white;
  margin: 40vh auto 0 auto;
  border-radius: 5px;

  h2 {
    font-weight: bold;
    margin-bottom: 3px;
  }

  p {
    font-size: 0.8rem;
  }

  input {
    display: block;
    border-radius: 5px;
    border: 1px solid $stroke-grey;
    height: 35px;
    width: 100%;
    padding: 2px 5px;
    color: $font-black;
    margin: 10px 0;
  }

  .actions {
    display: flex;
    justify-content: flex-end;
  }

  .cancel,
  .ok {
    display: block;
    border-radius: 3px;
    height: 25px;
    width: 80px;
    text-align: center;
  }

  .cancel {
    color: $stroke-grey;
  }

  .ok {
    color: $font-black;
  }
}
</style>
