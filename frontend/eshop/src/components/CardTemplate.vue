<template>
<h1>card view example</h1>
<div id="card-body">
    <img class="card-picture" :src="card.picturepath"/>
    <h3>{{ card.title }}</h3>
    <h4>{{ card.price }} ₽</h4>
    <h5>{{ card.description }}</h5>
</div>
</template>

<script>
import axios from "axios";
export default {
  name: 'CardTemplate',
  data() {
    return {
      card: {},
    };
  },

  methods: {
    getCard() {
    var bodyFormData = new FormData();
    bodyFormData.append('id', this.$route.params.id);
    axios({
      method: 'post',
      url: 'http://localhost:1111/api/card',
      data: bodyFormData
    })
    .then((response) => {
        this.card = response.data.message;
      });
    },
  },

  beforeMount() {
    this.getCard();
  }
};
</script>