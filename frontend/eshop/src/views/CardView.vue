<template>
  <div class="card">
    <CardTemplate/>
  </div>
</template>

<script>
// @ is an alias to /src
import CardTemplate from '@/components/CardTemplate.vue'
export default {
  name: 'CardView',
  components: {
    CardTemplate
  }
}
</script>