<template>
<HeaderTemplate/>
  <router-view/>
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
  color: #2c3e50;
}
</style>
