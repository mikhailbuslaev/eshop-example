<template>
  <PageUpButton/>
  <CatalogTemplate/>
  <CatalogPagination/>
</template>

<script>
// @ is an alias to /src
import PageUpButton from '@/components/PageUpButton.vue'
import CatalogTemplate from '@/components/CatalogTemplate.vue'
import CatalogPagination from '@/components/CatalogPagination.vue'
export default {
  name: 'CatalogView',
  components: {
    PageUpButton,
    CatalogTemplate,
    CatalogPagination
  }
}
</script>
