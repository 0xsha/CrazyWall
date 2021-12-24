<template>
  <div class="text-gray-400">
    <div v-if="$fetchState.pending">
      <Preloader/>
    </div>
    <p v-else-if="$fetchState.error">Error while fetching domain detail please retry later</p>
    <div v-else>


      <div class="flex items-center justify-center bg-draculaBackground py-4 mt-2 mb-2 rounded-lg ">

        <img class="h-14 w-14 mr-2" src="~/assets/images/browser.png" alt="">


        <h1 class="text-center text-3xl text-green-200 py-4">{{  domain }}</h1>

      </div>

      <div class="grid grid-cols-2 gap-4 py-4">
        <div class="py-4  border-1 border-purple-400 shadow-2xl rounded-lg dark:bg-gray-800">

          <div class="flex justify-center items-center">
            <img class="h-12 w-12 mr-2" src="~/assets/images/ip.png" alt="">

            <h1 class="text-xl text-center text-draculaGreen">IP Information</h1>
          </div>

          <ul class="px-12">
<!--           <li>{{this.ipData["count"]}}</li>-->
            <li v-if='this.ipData["ip"]' class="flex justify-between"> <span class="text-gray-50">IP </span> {{this.ipData["ip"]}}</li>
            <li v-if='this.ipData["country_name"]' class="flex justify-between"> <span class="text-gray-50">Country </span> {{this.ipData["country_name"]}}</li>
            <li v-if='this.ipData["currency"]["code"]' class="flex justify-between"> <span class="text-gray-50">Currency </span>{{this.ipData["currency"]["code"]}} |  {{this.ipData["currency"]["name"]}}  |  {{ this.ipData["currency"]["symbol"]}}</li>
            <li v-if='this.ipData["asn"]["asn"]' class="flex justify-between"> <span class="text-gray-50">ASN </span> {{this.ipData["asn"]["asn"]}}</li>
            <li v-if="this.ipData['country_code']" class="flex justify-between">  <span class="text-gray-50"> Flag</span>  <Flag class="rounded-lg" :code="this.ipData['country_code']"/> </li>
            <li v-if='this.ipData["latitude"]' class="flex justify-between"> <span class="text-gray-50">Latitude </span>  {{this.ipData["latitude"]}}</li>
            <li v-if='this.ipData["longitude"]' class="flex justify-between"> <span class="text-gray-50">Longitude </span>  {{this.ipData["longitude"]}}</li>

            <li class="flex justify-between items-center text-gray-50"> <span> Anonymous IP?</span> <span> <img v-if='this.ipData["threat"]["is_anonymous"] === false' src="~/assets/images/check.png" class="h-6 w-6"  alt="check"> <img v-if='this.ipData["threat"]["is_anonymous"]' src="~/assets/images/cross.png" class="h-12 w-12"  alt="check"> </span> </li>
            <li class="flex justify-between items-center text-gray-50"> <span> Bogon IP?</span> <span> <img v-if='this.ipData["threat"]["is_bogon"] === false' src="~/assets/images/check.png" class="h-6 w-6"  alt="check"> <img v-if='this.ipData["threat"]["is_anonymous"]' src="~/assets/images/cross.png" class="h-12 w-12"  alt="check"> </span> </li>
            <li class="flex justify-between items-center text-gray-50"> <span> Abuser IP?</span> <span> <img v-if='this.ipData["threat"]["is_known_abuser"] === false' src="~/assets/images/check.png" class="h-6 w-6"  alt="check"> <img v-if='this.ipData["threat"]["is_anonymous"]' src="~/assets/images/cross.png" class="h-12 w-12"  alt="check"> </span> </li>
            <li class="flex justify-between items-center text-gray-50"> <span> Attacker IP?</span> <span> <img v-if='this.ipData["threat"]["is_known_attacker"] === false' src="~/assets/images/check.png" class="h-6 w-6"  alt="check"> <img v-if='this.ipData["threat"]["is_anonymous"]' src="~/assets/images/cross.png" class="h-12 w-12"  alt="check"> </span> </li>
            <li class="flex justify-between items-center text-gray-50"> <span> Proxy ?</span> <span> <img v-if='this.ipData["threat"]["is_proxy"] === false' src="~/assets/images/check.png" class="h-6 w-6"  alt="check"> <img v-if='this.ipData["threat"]["is_anonymous"]' src="~/assets/images/cross.png" class="h-12 w-12"  alt="check"> </span> </li>
            <li class="flex justify-between items-center text-gray-50"> <span> Tor ?</span> <span> <img v-if='this.ipData["threat"]["is_tor"] === false' src="~/assets/images/check.png" class="h-6 w-6"  alt="check"> <img v-if='this.ipData["threat"]["is_anonymous"]' src="~/assets/images/cross.png" class="h-12 w-12"  alt="check"> </span> </li>
            <li class="flex justify-between items-center text-gray-50"> <span> Malicious ?</span> <span> <img v-if='this.ipData["threat"]["is_threat"] === false' src="~/assets/images/check.png" class="h-6 w-6"  alt="check"> <img v-if='this.ipData["threat"]["is_anonymous"]' src="~/assets/images/cross.png" class="h-12 w-12"  alt="check"> </span> </li>

          </ul>

        </div>

        <div class="py-4  border-1 border-purple-400 shadow-2xl rounded-lg dark:bg-gray-800">

          <div class="flex justify-center items-center">
            <img class="h-12 w-12 mr-2" src="~/assets/images/whois.png" alt="">

            <h1 class="text-xl text-center text-draculaGreen">Whois Information</h1>
          </div>

          <ul class="px-12">
            <li v-if='this.whois["WhoisRecord"]["contactEmail"]' class="flex justify-between"> <span class="text-gray-50">Contact email </span> {{this.whois["WhoisRecord"]["contactEmail"]}}</li>
            <li v-if='this.whois["WhoisRecord"]["estimatedDomainAge"]' class="flex justify-between"> <span class="text-gray-50">Estimated domain age </span> {{this.whois["WhoisRecord"]["estimatedDomainAge"]}}</li>
            <li v-if='this.whois["WhoisRecord"]["registrarName"]' class="flex justify-between"> <span class="text-gray-50">Registrar name </span> {{this.whois["WhoisRecord"]["registrarName"]}}</li>
            <li v-if='this.whois["WhoisRecord"]["registryData"]["createdDate"]' class="flex justify-between"> <span class="text-gray-50">Created date </span> {{this.whois["WhoisRecord"]["registryData"]["createdDate"]}}</li>
            <li v-if='this.whois["WhoisRecord"]["registryData"]["expiresDate"]' class="flex justify-between"> <span class="text-gray-50">Expire date </span> {{this.whois["WhoisRecord"]["registryData"]["expiresDate"]}}</li>

            <li v-if='this.whois["WhoisRecord"]["registryData"]["expiresDate"]' class="flex justify-between"> <span class="text-gray-50">Expire date </span> {{this.whois["WhoisRecord"]["registryData"]["expiresDate"]}}</li>

            <li v-if='this.whois["WhoisRecord"]["registryData"]["Status"]' class="flex justify-between"> <span class="text-gray-50">Status </span> {{this.whois["WhoisRecord"]["registryData"]["Status"]}}</li>

            <li v-if='this.whois["WhoisRecord"]["registryData"]["UpdatedDate"]' class="flex justify-between"> <span class="text-gray-50">Updated date </span> {{this.whois["WhoisRecord"]["registryData"]["UpdatedDate"]}}</li>

        <li v-if='this.whois["WhoisRecord"]["registryData"]["nameServers"]["hostNames"]' class="flex justify-between">
          <span class="text-gray-50">Host Names </span>

            <ul>
            <li v-for='hostname  in this.whois["WhoisRecord"]["registryData"]["nameServers"]["hostNames"]'>

              <span class="flex ">{{ hostname }}</span>

            </li>

            </ul>


        </li>

<!--            <li v-if='this.whois["WhoisRecord"]["registryData"]["nameServers"]["ips"]' class="flex justify-between"> <span class="text-gray-50">IPS</span> {{this.whois["WhoisRecord"]["registryData"]["nameServers"]["ips"]}}</li>-->




          </ul>

        </div>

      </div>




      <div class="grid grid-cols-2 gap-4 py-4">
        <div class="py-4  border-1 border-purple-400 shadow-2xl rounded-lg dark:bg-gray-800">

          <div class="flex justify-center items-center">
            <img class="h-12 w-12 mr-2" src="~/assets/icons/photo.svg" alt="">

            <h1 class="text-xl text-center text-draculaGreen">Mobile Shot</h1>
          </div>

          <img :src="screenshots[0]" class="rounded-lg py-2" alt="Desktop">
        </div>


        <div class="py-4  border-1 border-purple-400 shadow-2xl rounded-lg dark:bg-gray-800">

          <div class="flex justify-center items-center">
            <img class="h-12 w-12 mr-2" src="~/assets/icons/photo.svg" alt="">

            <h1 class="text-xl text-center text-draculaGreen">Desktop Shot</h1>
          </div>

          <img :src="screenshots[1]" class="rounded-lg py-2" alt="Desktop">
        </div>


        </div>

      <div class="grid grid-cols-1 gap-4 py-4">
        <div class="py-4  border-1 border-purple-400 shadow-2xl rounded-lg dark:bg-gray-800">

          <div v-if="history">
            <Association type="domain"  :search_id="this.search_id" :history="this.history"/>
          </div>

          <div v-else>
            <Association type="domain" :keyword="this.domain" :search_id="this.s_id"/>
          </div>

        </div>


      </div>



      <div v-if='this.ipData["latitude"]' class="grid grid-cols-1 gap-4 py-4">

        <Map :latitude='this.ipData["latitude"]' :longitude='this.ipData["longitude"]'></Map>

      </div>



    </div>

    </div>


</template>

<script>
import Association from "./Association"
import Preloader from "./Preloader";
import Flag from "./Flag";
export default {

  components : {Flag, Association,Preloader},

  props: ['domain', 'history', 'search_id'],


  data(){

    return {
      s_type : "domain",
      s_id: 0,
      whois: {},
      ipData :{},
      screenshots : {}
    }
  },

  async fetch() {


    if (this.history) {

      let details = await this.$axios.post("/api/v1/history/search", {
        s_type: this.s_type,
        search_id: this.search_id.toString()
      })
      this.ipData = details.data["ipData"]
      this.whois = details.data["whois"]
      this.s_id = details.data["searchID"]
      this.screenshots = details.data["screenShots"]

      return
    }



    if (!this.history) {

      let response =  await this.$axios.post("/api/v1/domain/search", {domain: this.domain})
      this.s_id = response.data["searchID"]
      this.whois = response.data["whois"]
      this.ipData = response.data["ipData"]
      this.screenshots = response.data["screenShots"]

    //  console.log(response.data["ipData"])
     // console.log(response.data["whois"])

    }

  }

}
</script>

<style scoped>

</style>
