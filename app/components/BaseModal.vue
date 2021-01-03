<template>
  <div v-show="show" class="base-modal-wrapper">
    <div class="base-modal">
      <div class="message">{{ message }}</div>
      <button class="close-btn" @click="close">OK</button>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from 'vue'
import { baseModalState } from '@/store'

export default Vue.extend({
  name: 'BaseModal',

  computed: {
    show(): boolean {
      return baseModalState.getShowModal
    },
    message(): string {
      return baseModalState.getMessage
    },
  },

  methods: {
    close(): void {
      baseModalState.setModal({ showModal: false, message: '' })
    },
  },
})
</script>

<style lang="scss" scoped>
@import 'assets/global';

.base-modal-wrapper {
  background-color: rgba(0, 0, 0, 0.8);
  position: absolute;
  height: 100vh;
  width: 100vw;
  top: 0;
  right: 0;
  padding: 0 10px;
  z-index: 10;
}

.message {
  margin: 0 0 10px 0;
}

.base-modal {
  padding: 5px;
  max-width: 450px;
  background-color: white;
  margin: 30vh auto 0 auto;
  text-align: center;
  border-radius: 5px;
}

.close-btn {
  display: block;
  margin: 0 auto;
  width: 100px;
  height: 35px;
  border-radius: 5px;
  background-color: $btn-blue;
  color: white;
}
</style>
