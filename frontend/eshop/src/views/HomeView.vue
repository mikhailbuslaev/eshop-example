<template>
  <div class="home">
    <CatalogBody/>
  </div>
</template>

<script>
// @ is an alias to /src
import CatalogBody from '@/components/CatalogBody.vue'

export default {
  name: 'HomeView',
  components: {
    CatalogBody
  }
}
</script>
