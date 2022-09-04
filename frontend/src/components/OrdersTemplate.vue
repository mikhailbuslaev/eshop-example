<template>
<div id="orders-container">
    <div v-for="order in orders" :key="order.id">
        <div id="orders-item">
            <h3>Order id: {{ order.cost }}</h3>
            <h3>Order cost: {{ order.cost }} ₽</h3>
            <h4>Address: {{ order.delivery.address }}</h4>
            <h4>Zip: {{ order.delivery.zip }}</h4>
            <h4>Receiver credentials: {{ order.delivery.receiver }}</h4>
            <div v-for="item in order.items" :key="item.id">
                <h4>Item id: {{ item.id }}</h4>
                <h4>Item quantity: {{ item.count }}</h4>
            </div>
        </div>
    </div>
</div>
</template>

<script>
import axios from "axios";
export default {
    name: 'OrdersTemplate',
    data() {
        return {
            orders: [],
            token: '',
            ordersLimit: 20,
            ordersBefore: 0
        };
    },
    methods: {
        loadOrders(startRow, quantity) {
            var bodyFormData = new FormData();
            bodyFormData.append('token', this.token);
            bodyFormData.append('quantity', quantity);
            bodyFormData.append('startrow', startRow);
            axios({
                method: 'post',
                url: 'http://localhost:1111/api/orders',
                data: bodyFormData
            })
            .then(response => {
                this.orders.push(...response.data.message);
            });
        }
    },
    beforeMount() {
        if (this.$cookies.isKey("token")) {
            this.token = this.$cookies.get("token");
            this.loadOrders(this.ordersBefore, this.ordersLimit);
        } else {
            this.$router.push({ path: '/login' });
        }
        this.ordersBefore = this.$route.params.page*this.ordersLimit;
    }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
#orders-container {
    display:flex;
    flex-direction:column;
    margin:10%;
    gap:20px 20px;
}

#orders-item {
    display:flex;
    flex-direction:column;
    align-items:flex-start;
    border-radius: 20px;
    border: 1px solid rgb(150,150,150);
    color:rgb(75,75,75);
}
</style>