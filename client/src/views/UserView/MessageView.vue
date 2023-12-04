<template>
  <el-table :data="user_message" border height="580px" show-empty>
    <el-table-column label="时间" prop="creat_at" width="280"/>
    <el-table-column label="商品序号" prop="id" width="200"/>
    <el-table-column label="商品名" prop="CommodityItem.item_name" width="200"/>
    <el-table-column label="当前价格" prop="CommodityItem.price" width="200"/>
    <el-table-column label="商家" prop="CommodityItem.Seller.username" width="200"/>
    <el-table-column label="平台" prop="CommodityItem.Platform.name" width="170"/>
  </el-table>
</template>

<script>

export default {
  name: "user_message",
  data() {
    return {
      user_message: localStorage.getItem('user_message')
        ? JSON.parse(localStorage.getItem('user_message'))
        : [],
    };
  },
  created() {
    this.findAll()
  },
  methods: {
    findAll(){
      this.request.get("/message/all").then((res) => {
        if (res.status === 200) {
          localStorage.setItem("user_message", JSON.stringify(res.data));
        } else {
          this.$message.error(res.message);
        }
      });
    }
  }
};

</script>