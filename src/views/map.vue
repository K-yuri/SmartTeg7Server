<template>
  <GoogleMap
  api-key="AIzaSyDDXDdoYwn-bRRX0LseKP0C0EpRnIxa0M4"
  style="width: 100%; height: 500px"
  :center="center"
  :zoom="15"
  >
    <Marker :options="{ position: center }" />
  </GoogleMap>
</template>

<script>
import { defineComponent } from 'vue'
import { GoogleMap, Marker } from 'vue3-google-map'

export default defineComponent({
  components: { GoogleMap, Marker },
  setup() {
    const center = { lat:36.322344, lng:127.338010 }

    return { center }
  },
})
</script>
