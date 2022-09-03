<template>
<div id="shoppingcart-body">
<h1>your shopping cart:</h1>
    <div v-for="[cartItemId, cartItem] in cartItems" :key="cartItemId">
        <div class="shoppingcart-item">
            <img id="cartitem-picture" :src="cartItem.cardData.picturepath"/>
            <h4 class="cartitem-text">{{ cartItem.cardData.title }}</h4>
            <h4 class="cartitem-text">{{ cartItem.cardData.price }}*{{ cartItem.itemData.count }} = {{this.renderItemCost(cartItem.cardData.price, cartItem.itemData.count)}}₽</h4>
            <div id="removeitem-button-container">
                <button id="removeitem-button" v-on:click="deleteItem(cartItem.itemData.id)">
                <h4>&#215;</h4>
                </button>
            </div>
        </div>
    </div>
<h3>Summary cost: {{ showSummaryCost }}₽</h3>
<button id="createorder-button" v-on:click="showDeliveryForm" v-if="deliveryFormHidden">Make order</button>
<div id="deliveryForm-container" v-if="!deliveryFormHidden">
    <form id="deliveryform" @submit="createOrder">
        <label for="fname">
            <h3>Zip code</h3>
        </label>
        <input type="text" v-model="deliveryData.zip">
        <label for="lname">
            <h3>Address</h3>
        </label>
        <input type="text" v-model="deliveryData.address">
        <label for="lname">
            <h3>Receiver name</h3>
        </label>
        <input type="text" v-model="deliveryData.receiver">
        <input type="submit" value="Create order">
    </form>
</div>
</div>

</template>

<script>
import axios from "axios";
export default {
    name: 'ShoppingCartTemplate',

    data() {
        return {
        cartId: 0,
        token: '',
        cartItems: new Map(),
        emptyCard: {id:'', title:'', price:0, picturepath:'', count:''},
        userId: 0,
        summaryCost: 0,
        deliveryData: {zip:'', address:'', receiver:''},
        deliveryFormHidden: true
        };
    },

    methods: {
    getShoppingCart() {
        var bodyFormData = new FormData();
        bodyFormData.append('token', this.token);
        axios({
        method: 'post',
        url: 'http://localhost:1111/api/shopping_cart',
        data: bodyFormData
        })
        .then((response) => {
            response.data.message.items.forEach(element => {
                this.cartItems.set(element.id, {itemData:element, cardData:this.emptyCard});
                this.getCard(element.id, element.cardid);
            });
        });
      },
    getCard(cartId = '', cardId = '') {
        var bodyFormData = new FormData();
        bodyFormData.append('id', cardId);
        axios({
        method: 'post',
        url: 'http://localhost:1111/api/card',
        data: bodyFormData
        })
        .then((response) => {
            let bufObj = this.cartItems.get(cartId);
            this.cartItems.set(bufObj.itemData.id, {itemData:bufObj.itemData, cardData:response.data.message});
            let obj = this.cartItems.get(cartId);
            this.increaseSummaryCost(this.renderItemCost(obj.itemData.count, obj.cardData.price));
        });
      },
    renderItemCost(first='1', second='1') {
        return parseInt(first)*parseInt(second);
      },
    increaseSummaryCost(itemCost = 0) {
        this.summaryCost += itemCost;
      },
    calculateSummaryCost() {
        this.summaryCost = 0;
        this.cartItems.forEach((value) => {
            this.increaseSummaryCost(this.renderItemCost(value.itemData.count, value.cardData.price));
        });
      },
    deleteItem(cartItemId = '') {
        var bodyFormData = new FormData();
        bodyFormData.append('token', this.token);
        bodyFormData.append('cartitemid', cartItemId);
        axios({
        method: 'post',
        url: 'http://localhost:1111/api/shopping_cart/delete_item',
        data: bodyFormData
        })
        .then(
            this.cartItems.delete(cartItemId),
            this.calculateSummaryCost()
        );
      },
    createOrder() {
        var bodyFormData = new FormData();
        bodyFormData.append('token', this.token);
        bodyFormData.append('delivery-zip', this.deliveryData.zip);
        bodyFormData.append('delivery-address', this.deliveryData.address);
        bodyFormData.append('delivery-receiver', this.deliveryData.receiver);
        axios({
        method: 'post',
        url: 'http://localhost:1111/api/orders/add',
        data: bodyFormData
        })
        .then(
            alert('Order successfully created'),
            this.$router.push({ path: '/user/orders/0' })
        );
      },
    showDeliveryForm() {
        this.deliveryFormHidden = false;
    }
    },

    beforeMount() {
        this.summaryCost = 0;
        if (this.$cookies.isKey("token") && this.$cookies.isKey("userId")) {
            this.userId = this.$cookies.get("userId");
            this.token = this.$cookies.get("token");
            this.getShoppingCart();
        } else {
            this.$router.push({ path: '/login' });
        }
    },

    computed: {
        showSummaryCost: function() {
            return this.summaryCost
        }
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.shoppingcart-item {
    display: flex;
    flex-direction: row;
    gap: 15%;
    justify-content:center;
}

#removeitem-button {
    background-color:white;
    border:0px;
    font-size:30px;
    color:rgb(75, 75, 75);
}

#removeitem-button:hover {
    color: red;
}

.cartitem-text {
    padding-top:25px;
    color:rgb(100, 100, 100);
}

#cartitem-picture {
    width:100px;
}
</style>