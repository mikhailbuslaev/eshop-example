<template>
<div class="carousel-body">
    <div v-for="card in cards" :key="card.id">
    <a v-bind:href="'http://localhost:8081/#/catalog/'+card.id">
      <div class="card">
        <img class="card-picture" :src="card.picturepath"/>
        <h3>{{ card.title }}</h3>
        <h4>{{ card.price }} ₽</h4>
      </div>
      </a>
    </div>
</div>
<button class="btn btn-next">&#11166;</button>
<button class="btn btn-prev">&#11164;</button>
</template>

<script>
import axios from "axios";
export default {
    name: 'CarouselTemplate',

    data() {
        return {
        cards: [],
        };
    },

    methods: {
    getCards() {
        var bodyFormData = new FormData();
        bodyFormData.append('quantity', 15);
        bodyFormData.append('startrow', 0);
        axios({
        method: 'post',
        url: 'http://localhost:1111/api/cards',
        data: bodyFormData
        })
        .then((response) => {
            this.cards = response.data.message;
        });
        }
    },
    beforeMount() {
        this.getCards();
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
.carousel-body {
  margin-left:5%;
  width:90%;
  overflow: hidden;
  display: flex;
  flex-wrap: nowrap;
  gap: 20px 20px;
}

.card {
  width: 200px;
  background: #ccc;
  margin: 1em auto;
}

.card-picture {
  width: 200px;
}

.btn {
  position: absolute;
  width: 40px;
  height: 40px;
  padding: 10px;
  border: none;
  border-radius: 50%;
  z-index: 10px;
  cursor: pointer;
  color:white;
  background-color: #ccc;
  font-size: 18px;
}

.btn:active {
  transform: scale(1.1);
}

.btn-prev {
  top: 170%;
  left: 2%;
}

.btn-next {
  top: 170%;
  right: 2%;
}
</style>