<template>
<div id="carousel-wrapper" ref="carouselWrapper">
  <div id="carousel-body" ref="carouselBody">
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
  <button class="btn btn-next" v-on:click="moveToLeft">&#8250;</button>
  <button class="btn btn-prev" v-on:click="moveToRight">&#8249;</button>
</div>
</template>

<script>
import axios from "axios";
export default {
    name: 'CarouselTemplate',

    data() {
        return {
        cards: [],
        counter : 0,
        clicksLimit : 0
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
      },

      moveCarousel() {
        this.$refs.carouselBody.style.transform = 'translateX('+(200*this.counter)+'px';
      },

      moveToLeft : function(event) {
        if (event) {
          if (!(this.counter===this.clicksLimit)) {
            this.counter--;
            this.moveCarousel();
          }
        }
      },

      moveToRight : function(event) {
        if (event) {
          if (!(this.counter===0)) {
            this.counter++;
            this.moveCarousel();
          }
        }
      }
    },

    beforeMount() {
      this.getCards();
    },

    mounted() {
      this.clicksLimit = -16 + Math.floor(this.$refs.carouselWrapper.clientWidth/200);
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#carousel-wrapper {
  overflow: hidden;
  position: relative;
  margin-left:5%;
  width:90%;
}

#carousel-body {
  transition: transform 0.3s ease-in-out;
  display: flex;
  flex-wrap: nowrap;
  gap: 20px 20px;
}

.card {
  width: 200px;
  background: #ccc;
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
  top: 50%;
  left: 2%;
}

.btn-next {
  top: 50%;
  right: 2%;
}

a:link {
  text-decoration: none;
  color:#2c3e50;
}

a:visited {
  text-decoration: none;
  color:#2c3e50;
}
</style>