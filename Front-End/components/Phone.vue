<template>
  <div class="text-gray-400">
    <div v-if="$fetchState.pending">
      <Preloader/>
    </div>
    <p v-else-if="$fetchState.error">Error while fetching number detail please retry later</p>
    <div v-else>


      <div class="flex items-center justify-center bg-draculaBackground py-4 mt-2 mb-2  rounded-lg">

        <img class="h-14 w-14 mr-2" src="~/assets/icons/hearing.svg" alt="">


        <h1 class="text-center text-3xl text-green-200 py-4">+{{  phone }}</h1>

      </div>
      <div class="grid grid-cols-2 gap-4">


        <div class="py-4  border-1 border-purple-400 shadow-2xl rounded-lg dark:bg-gray-800">


          <div class="flex justify-center items-center">
            <img class="h-14 w-14 mr-2" src="~/assets/icons/loupe.svg">

            <h1 class="text-2xl text-draculaGreen text-center">Google Dorks</h1>
          </div>

          <accordion title="Social Media">
            <div v-for="links in googleDetails.result['socialMedia']">

              <a class="px-6" :href="links.URL" target="_blank">{{ cleanLinks(links.URL) }}</a>
            </div>

          </accordion>

          <accordion title="Disposable Providers">
            <div v-for="links in googleDetails.result['disposableProviders']">
              <a class="px-6" :href="links.URL" target="_blank">{{ cleanLinks(links.URL) }}</a>

            </div>
          </accordion>


          <accordion title="Reputation">
            <div v-for="links in googleDetails.result['reputation']">
              <a class="px-6" :href="links.URL" target="_blank">{{ cleanLinks(links.URL) }}</a>

            </div>
          </accordion>


          <accordion title="Reputation">
            <div v-for="links in googleDetails.result['individuals']">
              <a class="px-6" :href="links.URL" target="_blank">{{ cleanLinks(links.URL) }}</a>

            </div>
          </accordion>

          <accordion title="General">
            <div v-for="links in googleDetails.result['general']">
              <a class="px-6" :href="links.URL" target="_blank">{{ cleanLinks(links.URL) }}</a>

            </div>
          </accordion>


        </div>


        <div class="py-4  border-1 border-purple-400 shadow-2xl rounded-lg dark:bg-gray-800">

          <div class="flex justify-center items-center">
            <img class="h-14 w-14 mr-2" src="~/assets/icons/note.svg" alt="">

            <h1 class="text-2xl text-draculaGreen text-center">Number information</h1>
          </div>

          <ul class="px-12 py-2">
            <li class="flex justify-between"> <span class="text-gray-50"> Valid: </span>{{ numDetails.result["valid"] }}</li>

            <li class="flex justify-between" v-if="numDetails.result['valid']"> <span class="text-gray-50">Country: </span> {{ numDetails.result["country_name"] }}</li>
            <li  v-if="numDetails.result['valid']" class="flex justify-between items-center py-2">
              <span class="text-gray-50">Flag: </span>
              <Flag class="flex justify-between" :code="numDetails.result['country_code']"/>
            </li>
            <li class="flex justify-between" v-if="numDetails.result['valid']"> <span class="text-gray-50"> Prefix: </span> {{ numDetails.result["country_prefix"] }}</li>
            <li class="flex justify-between" v-if="numDetails.result['valid']"> <span class="text-gray-50"> Code: </span> {{ numDetails.result["country_code"] }}</li>


            <li class="flex justify-between"  v-if="numDetails.result['location']"> <span class="text-gray-50"> Location: </span> {{ numDetails.result["location"] }}</li>
            <li class="flex justify-between"  v-if="numDetails.result['carrier']"> <span class="text-gray-50"> Carrier: </span> {{ numDetails.result["carrier"] }}</li>
            <li class="flex justify-between"  v-if="numDetails.result['valid']">  <span class="text-gray-50"> Line Type: </span> {{ numDetails.result["line_type"] }}</li>
          </ul>
        </div>

      </div>

      <div class="grid grid-cols-1 gap-4 py-4">
        <div class="py-4  border-1 border-purple-400 shadow-2xl rounded-lg dark:bg-gray-800">

          <div v-if="history">
            <Association type="number"  :search_id="this.search_id" :history="this.history"/>
          </div>

          <div v-else>
            <Association type="number" :keyword="this.phone" :search_id="this.s_id"/>
          </div>

        </div>


      </div>
    </div>
  </div>


</template>

<style>
</style>
<script>
import Accordion from "./Accordion";
import Preloader from "./Preloader";
import Flag from "./Flag";
import Association from "./Association";

export default {
  components: {Association, Flag, Preloader, Accordion},

  // loadingIndicator: {
  //   name: 'circle',
  //   color: '#3B8070',
  //   background: 'white'
  // },


  props: ['phone', 'history', 'search_id'],

  data() {
    return {
      numDetails: [],
      googleDetails: [],
      link: "",
      s_type: "number",
      s_id : 0
    }

  },

  computed: {},


  methods: {

    cleanLinks(val) {

      let pattern = /site%3A(.*?)\+/g

      //console.log(escape(unescape(val)))

      let matches = pattern.exec(escape(unescape(val)));

      if (matches) {

        val = matches[1]

        return val
      }

      // well for now
      val = "Dork"
      //console.log(this.link)
      return val
    }
  },
  async fetch() {


    if (this.history) {

      let details = await this.$axios.post("/api/v1/history/search", {
        s_type: this.s_type,
        search_id: this.search_id.toString()
      })
      this.numDetails = details.data["numverify"]
      this.googleDetails = details.data["google"]
      this.searchID = details.data["searchID"]


      return
    }

    let details = await this.$axios.post("/api/v1/phone/search", {phone: this.phone})
    this.numDetails = details.data["numverify"]
    this.googleDetails = details.data["google"]
    this.s_id = details.data["searchID"]
    //this.$emit('art' , true)

  }
}
</script>
