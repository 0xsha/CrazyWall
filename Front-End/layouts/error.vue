<template>
  <div>
    <div v-if="error.statusCode === 404" class="flex justify-center mx-auto min-h-screen">
      <img @click="$router.push('/')"  class="cursor-pointer w-3/5 rounded-lg my-4" loading="lazy" src="~/assets/images/404.webp" alt="">
    </div>
    <h1 v-else>An error occurred</h1>
    <NuxtLink to="/"></NuxtLink>
  </div>
</template>

<script>
export default {
  props: ['error'],
  layout:'auth'
}
</script>

<style scoped>

</style>
