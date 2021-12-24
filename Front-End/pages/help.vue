<template>
  <div class="container text-draculaForeground px-4">

    <h1 class="text-2xl  font-bold py-4"> Frequently asked questions</h1>
<!--    <h3 class="py-2 text-xl"> Can't find the answer you are looking for? Don't hesitate to reach out to us.</h3>-->
    <hr class="border-1 px-4 border-gray-400 pb-2">
    <div class="grid  grid-col-1 md:grid-cols-2 gap-2 py-4">


      <div>
        <p class="text-xl py-2 "> What is this? </p>
        <p class="text-gray-400">
          Think of CrazyWall as a digital evidence board. It accepts and treats different user inputs, namely IP address, username, full name, domain name. after detecting and processing your input based on your input, it returns different data.
        </p>
      </div>

      <div>
        <p class="text-xl py-2"> For whom? </p>
        <p class="text-gray-400">
          We created CrazyMap for everyone , Individuals can find their privacy breaches. Journalists, private investigators, fraud analysts, security and malware researchers (e.g., to find the guy behind APT), red-teamers can use it to create an excellent pretext, an organization to monitor their sentiment, and breach monitoring. Threat intel platforms can show the results to their feed or pass it to machine learning algorithms.
        </p>
      </div>

    </div>
    <div class="grid grid-col-1 md:grid-cols-2 gap-2 py-4">


      <div>
        <p class="text-xl py-2"> How to use it?</p>
        <p class="text-gray-400">
          All you have to do is go to the main page, type your keyword, select your mode between databases and live and hit the enter button.
          Valid inputs are :
          a domain name, e.g., <span class="bg-black py-1  px-2 text-sm text-white rounded-lg">domain.com</span>
          a phone number (international format), e.g.,
          <span class="bg-black py-1  px-2 text-sm text-white rounded-lg">00123456789</span>
          a full name, e.g.,  <span class="bg-black py-1  px-2 text-sm text-white rounded-lg">john due</span>
          a user name, e.g.,  <span class="bg-black py-1  px-2 text-sm text-white rounded-lg">sherlock</span>
          an IP address, e.g.,
          <span class="bg-black py-1  px-2 text-sm text-white rounded-lg">1.2.3.4</span>
          a keyword for live search, e.g.,

          <span class="bg-black py-1  px-2 text-sm text-white rounded-lg">crypto</span>
        <p>



        </p>

      </div>

      <div>
        <p class="text-xl py-2">What data do you collect?</p>
        <p class="text-gray-400">
          Besides your data on registration and analytics data, we do not collect anything else. We do not share, sell, or distribute your data.
        </p>
      </div>


    </div>


    <div class="grid grid-col-1 md:grid-cols-2 gap-2 py-4">


      <div>
        <p class="text-xl py-2"> Is it legal? </p>
        <p class="text-gray-400">
          CrazyWall does not collect or store any breaches; We use 3rd sources and scaping methods to conduct our results; either we like it or not, the data is available here or there, and we can't just ignore it.
        </p>
      </div>

      <div>
        <p class="text-xl py-2"> Is it free? </p>
        <p class="text-gray-400">
          For independent journalists and researchers, it's free and always will be. If you would like to use our engine in your commercial product or organization, please <a class="bg-black py-1 text-white px-2 rounded-lg" href="mailto:hello@crazywall.io" target="_blank"> contact </a> us.
        </p>
      </div>

    </div>
    <div class="grid grid-col-1 md:grid-cols-2 gap-2 py-4">


      <div>
        <p class="text-xl py-2"> Is there a limitation? </p>
        <p class="text-gray-400">
          You are allowed to search 30 entities, and entities are limited to 1k. We also remove credentials from views. Don't hesitate to <a class="bg-black py-1 text-white px-2 rounded-lg" href="mailto:hello@crazywall.io" target="_blank"> contact </a> us if you need to lift limitations, complete raw data, monitoring, or API access; please do not hesitate to <a class="bg-black py-1 text-white px-2 rounded-lg" href="mailto:hello@crazywall.io" target="_blank"> contact </a> us.
        </p>
      </div>

      <div>
        <p class="text-xl py-2">Is there a way to monitor mentions about me or my organization?</p>
        <p class="text-gray-400">
          Yes, you can use our monitoring API to get notifications about new mentions in your favorite channel. e.g., (slack, webhook, e-mail)        </p>
      </div>

    </div>
    <div class="grid grid-col-1 md:grid-cols-2 gap-2 py-4">


      <div>
        <p class="text-xl py-2"> Can I delete mentions about my organization or me?</p>
        <p class="text-gray-400">
          As mentioned, CrazyMap is not the source of any of the data returned by dashboard or API, but you can <a class="bg-black py-1 text-white px-2 rounded-lg" href="mailto:hello@crazywall.io" target="_blank"> contact </a> us to get information about the source of detailed data.         </p>
      </div>


      <div>
        <p class="text-xl py-2">Can't find the answer you are looking for?</p>
        <p class="text-gray-400">
          I still can not find the answer I'm looking for, Don't hesitate to <a class="bg-black py-1 text-white px-2 rounded-lg" href="mailto:hello@crazywall.io" target="_blank"> reach </a> out to us.
        </p>
      </div>



    </div>


  </div>

</template>

<script>
export default {
  name: "help"
}
</script>

<style scoped>

</style>
