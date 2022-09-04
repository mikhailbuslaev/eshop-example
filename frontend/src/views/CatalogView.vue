<template>
  <PageUpButton/>
  <CatalogNavigation/>
  <CatalogTemplate/>
  <CatalogPagination/>
</template>

<script>
// @ is an alias to /src
import PageUpButton from '@/components/PageUpButton.vue'
import CatalogTemplate from '@/components/CatalogTemplate.vue'
import CatalogPagination from '@/components/CatalogPagination.vue'
import CatalogNavigation from '@/components/CatalogNavigation.vue'
export default {
  name: 'CatalogView',
  components: {
    PageUpButton,
    CatalogTemplate,
    CatalogPagination,
    CatalogNavigation
  }
}
</script>
