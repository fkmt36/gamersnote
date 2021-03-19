<template>
  <div class="container">
    <h1 v-if="error.statusCode === 404">Page not found</h1>
    <h1 v-else-if="error.statusCode === 403">Forbidden</h1>
    <h1 v-else>An error occurred</h1>
    <p>{{ error.message }}</p>
    <NuxtLink to="/">トップページに戻る</NuxtLink>
  </div>
</template>

<script>
export default {
  props: ['error'],
}
</script>
