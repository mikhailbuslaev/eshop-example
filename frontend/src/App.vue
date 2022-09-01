<template>
<HeaderTemplate/>
  <router-view :key="$route.fullPath"/>
<FooterTemplate/>
</template>

<script>
import HeaderTemplate from '@/components/HeaderTemplate.vue'
import FooterTemplate from '@/components/FooterTemplate.vue'
export default {
  components: {
    HeaderTemplate,
    FooterTemplate
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: rgb(75, 75, 75);
}
</style>
