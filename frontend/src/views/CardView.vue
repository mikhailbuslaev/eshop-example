<template>
    <CardTemplate/>
    <CarouselTemplate/>
</template>

<script>
// @ is an alias to /src
import CardTemplate from '@/components/CardTemplate.vue'
import CarouselTemplate from '@/components/CarouselTemplate.vue'
export default {
  name: 'CardView',
  components: {
    CardTemplate,
    CarouselTemplate
  }
}
</script>