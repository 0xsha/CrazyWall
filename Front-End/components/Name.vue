<template>
  <div class="text-gray-400">
    <div v-if="$fetchState.pending">
      <Preloader/>
    </div>
    <p v-else-if="$fetchState.error">Error while fetching email detail please retry later</p>
    <div v-else>


      <div class="flex items-center justify-center bg-draculaBackground py-4 mt-2  rounded-lg">

        <img class="h-14 w-14 mr-2" src="~/assets/icons/passport.svg" alt="">


        <h1 class="text-center text-3xl text-green-200 py-4">{{  name }}</h1>

      </div>

      <div class="grid grid-cols-1 gap-4 py-4">
        <div class="py-4  border-1 border-purple-400 shadow-2xl rounded-lg dark:bg-gray-800">

          <div v-if="history">
            <Association type="name"  :search_id="this.search_id" :history="this.history"/>
          </div>

          <div v-else>
            <Association type="name" :keyword="this.name" :search_id="this.s_id"/>
          </div>

        </div>


      </div>
    </div>
  </div>

</template>

<script>
import Association from "./Association"
import Preloader from "./Preloader";
export default {

  components : {Association,Preloader},

  props: ['name', 'history', 'search_id'],


  data(){

    return {
      s_type : "name",
      s_id: 0
    }
  },

  async fetch() {


    if (!this.history) {

      let response =  await this.$axios.post("/api/v1/name/search", {name: this.name})
      this.s_id = response.data["searchID"]
    }

  }

}
</script>

<style scoped>

</style>
