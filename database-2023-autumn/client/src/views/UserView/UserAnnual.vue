<template>
  <el-carousel :interval="4000" type="card" height="400px" style="margin-top: 50px">
    <el-carousel-item style="border-radius: 15px">
      <h2 text="2xl" justify="center" style="margin-top: 40px">收藏最多的商品种类:</h2>
      <h3 text="2xl" justify="center" style="margin-left: 100px; margin-top: 30px">名称:{{this.annual_data.category}}</h3>
      <h3 text="2xl" justify="center" style="margin-left: 100px; margin-top: 20px">收藏次数:{{this.annual_data.category_num}}</h3>
    </el-carousel-item>
    <el-carousel-item style="border-radius: 15px">
      <h2 text="2xl" justify="center" style="margin-top: 40px">收藏最多商品:</h2>
      <h3 text="2xl" justify="center" style="margin-left: 100px; margin-top: 30px">名称:{{this.annual_data.Commodity.default_name}}</h3>
      <h3 text="2xl" justify="center" style="margin-left: 100px; margin-top: 20px">种类:{{ this.annual_data.Commodity.category}}</h3>
      <h3 text="2xl" justify="center" style="margin-left: 100px; margin-top: 20px">产地:{{ this.annual_data.Commodity.produce_address}}</h3>
      <h3 text="2xl" justify="center" style="margin-left: 100px; margin-top: 20px">生产时间:{{ this.annual_data.Commodity.produce_at}}</h3>
      <h3 text="2xl" justify="center" style="margin-left: 100px; margin-top: 20px">收藏次数:{{this.annual_data.commodity_num}}</h3>
    </el-carousel-item>
    <el-carousel-item style="border-radius: 15px">
      <h2 text="2xl" justify="center" style="margin-top: 40px">收藏最多商家:</h2>
      <h3 text="2xl" justify="center" style="margin-left: 100px; margin-top: 30px">商家名:{{this.annual_data.Seller.username}}</h3>
      <h3 text="2xl" justify="center" style="margin-left: 100px; margin-top: 20px">邮箱:{{ this.annual_data.Seller.email}}</h3>
      <h3 text="2xl" justify="center" style="margin-left: 100px; margin-top: 20px">地址:{{ this.annual_data.Seller.address}}</h3>
      <h3 text="2xl" justify="center" style="margin-left: 100px; margin-top: 20px">收藏次数:{{this.annual_data.seller_num}}</h3>
    </el-carousel-item>
    <el-carousel-item style="border-radius: 15px">
      <h2 text="2xl" justify="center" style="margin-top: 40px">收藏最多平台:</h2>
      <h3 text="2xl" justify="center" style="margin-left: 100px; margin-top: 30px">平台名:{{this.annual_data.Platform.name}}</h3>
      <h3 text="2xl" justify="center" style="margin-left: 100px; margin-top: 20px">国家:{{ this.annual_data.Platform.country}}</h3>
      <h3 text="2xl" justify="center" style="margin-left: 100px; margin-top: 20px">URL:{{ this.annual_data.Platform.url}}</h3>
      <h3 text="2xl" justify="center" style="margin-left: 100px; margin-top: 20px">收藏次数:{{this.annual_data.platform_num}}</h3>
    </el-carousel-item>
    <el-carousel-item style="border-radius: 15px">
      <h1 text="2xl" justify="center" style="margin-top: 40px">累计收藏总数:</h1>
      <span style="margin-left: 100px; font-size: 150px; color: #79bbff">{{this.annual_data.favorite_num}}</span>
    </el-carousel-item>
    <el-carousel-item style="border-radius: 15px">
      <h1 text="2xl" justify="center" style="margin-top: 40px">累计接收消息: </h1>
      <span style="margin-left: 100px; font-size: 150px; color: #79bbff">{{this.annual_data.message_num}}</span>
    </el-carousel-item>
    <el-carousel-item style="border-radius: 15px">
      <h1 text="2xl" justify="center" style="margin-top: 40px">收藏商品价格变动次数:</h1>
      <span style="margin-left: 100px; font-size: 150px; color: #79bbff">{{this.annual_data.price_change_num}}</span>
    </el-carousel-item>
  </el-carousel>
</template>

<script>
export default {
  name: "user_annual",
  data() {
    return {
      annual_data:{
        category: null,
        category_num: 0,
        Commodity: {
          category: null,
          default_name: null,
          id: 0,
          produce_address: null,
          produce_at: null,
        },
        commodity_num: 0,
        favorite_num: 0,
        message_num: 0,
        Platform: {
          country: null,
          id: 0,
          name: null,
          url: null
        },
        platform_num: 0,
        price_change_num: 0,
        Seller: {
          address: null,
          email: null,
          id: 0,
          password: null,
          username: null,
        },
        seller_num: 0,
      }
    };
  },
  created() {
    this.find_data();
  },
  methods: {
    find_data(){
      this.request.get("/annual/summary").then((res) => {
        if (res.status === 200) {
          this.annual_data=res.data
        }
      }).catch(error => {
        this.$message.error(error.response.data.message);
      });
    }
  }
};
</script>

<style scoped>
.el-carousel__item h1 {
  color: rgba(152, 8, 68, 0.8);
  margin-left: 50px;
}

.el-carousel__item h2 {
  color: rgba(99, 109, 220, 0.89);
  margin-left: 50px;
}

.el-carousel__item h3 {
  color: #475669;
  opacity: 0.75;
  margin-left: 50px;
}

.el-carousel__item:nth-child(2n) {
  background-color: #99a6bf;
}

.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dde6;
}
</style>
