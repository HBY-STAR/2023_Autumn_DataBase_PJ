<template>
  <div style="margin-top: 10px;margin-left: 10px;">
    <div style="width: 300px; float: left">
      <el-card style="width: 300px;">
        <el-form label-width="80px" size="small" style="margin-left: 10px">
          <div style="height: 30px">
            <header style="text-align: center; color: #337ecc">
              商家信息
            </header>
          </div>
          <el-form-item label="用户ID:">
            <span>{{ seller_data.id }}</span>
          </el-form-item>
          <el-form-item label="用户名:">
          <span>
            {{ seller_data.username }}
          </span>
          </el-form-item>
<!--          <el-form-item label="密码:">-->
<!--          <span>-->
<!--            {{ seller_data.password }}-->
<!--          </span>-->
<!--          </el-form-item>-->
          <el-form-item label="邮箱:">
          <span>
            {{ seller_data.email }}
          </span>
          </el-form-item>
          <el-form-item label="地址:">
          <span>
            {{ seller_data.address }}
          </span>
          </el-form-item>
          <el-form-item label="类型:">
            <span>{{ type }}</span>
          </el-form-item>
        </el-form>
      </el-card>
    </div>
  </div>
</template>

<script>

export default {
  name: "seller_data",
  data() {
    return {
      seller_data: localStorage.getItem('seller_data')
        ? JSON.parse(localStorage.getItem('seller_data'))
        : {},
      type: localStorage.getItem('user_type')
        ? JSON.parse(localStorage.getItem('user_type'))
        : '未知',
    }
  },
  created() {
    if(localStorage.getItem("seller_data") == null){
      this.findCurSeller()
    }
  },
  methods: {
    findCurSeller() {
      this.request.get("/sellers/data").then((res) => {
        if (res.status === 200) {
          localStorage.setItem("seller_data", JSON.stringify(res.data));
        } else {
          this.$message.error(res.message);
        }
      })
    }
  },

}
</script>

