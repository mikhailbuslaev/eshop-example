<template>
  <CatalogTemplate/>
  <CatalogPagination/>
</template>

<script>
// @ is an alias to /src
import CatalogTemplate from '@/components/CatalogTemplate.vue'
import CatalogPagination from '@/components/CatalogPagination.vue'
export default {
  name: 'CatalogView',
  components: {
    CatalogTemplate,
    CatalogPagination
  }
}
</script>
