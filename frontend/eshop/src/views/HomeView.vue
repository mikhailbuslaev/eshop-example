<template>
  <div class="home">
    <CatalogTemplate/>
  </div>
</template>

<script>
// @ is an alias to /src
import CatalogTemplate from '@/components/CatalogTemplate.vue'
export default {
  name: 'HomeView',
  components: {
    CatalogTemplate
  }
}
</script>
