<template>
<div id="map" class="w-full h-96 rounded-lg"></div>
</template>

<script>
import mapboxgl from 'mapbox-gl'
import 'mapbox-gl/dist/mapbox-gl.css';


export default {

  props:["longitude" , "latitude"],

  mounted(){

    mapboxgl.accessToken = this.$config.keys.mapBox;
    const map = new mapboxgl.Map({
      container: 'map', // container ID
      style: 'mapbox://styles/mapbox/dark-v10', // style URL
      center: [ this.longitude , this.latitude], // starting position [lng, lat]
      zoom: 12  // starting zoom
    });


  },

  name: "Map"
}
</script>

<style>

.mapboxgl-ctrl-attrib-inner {
  visibility: hidden;
}

</style>
