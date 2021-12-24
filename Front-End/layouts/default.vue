<template>
  <div class="dark:bg-gray-900">
    <Navbar />
    <Nuxt />
    <Footer />
  </div>
</template>

<style>

/* dark magic to override default vue-toast behavior */
.toasted {
  @apply rounded-lg h-20 w-60 !important;
}
.action {
  @apply flex justify-center items-center rounded-full w-8 h-8 text-center text-xl    !important;
}

html {
  font-family: 'Inter', sans-serif;

}
.container {
  @apply  min-h-screen mx-auto
}

</style>
<script>
import Navbar from "../components/Navbar";

import Footer from "../components/Footer";
export default {
  components: {Footer , Navbar}

// components: navbar,
  // data(){
  // return {
  //   auth : this.$auth.loggedIn
  // }
  // }

}
</script>
