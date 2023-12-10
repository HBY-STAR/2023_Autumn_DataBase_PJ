<template>
  <div style="margin-top: 10px;margin-left: 10px;">
    <div style="width: 300px; float: left">
      <el-card style="width: 300px;">
        <el-form label-width="80px" size="small" style="margin-left: 10px">
          <div style="height: 30px"><header style="text-align: center; color: #337ecc">个人信息</header></div>
          <el-form-item label="用户ID:">
            <span>{{ user_data.id }}</span>
          </el-form-item>
          <el-form-item label="用户名:">
            <span>{{ user_data.username }}</span>
          </el-form-item>
<!--          <el-form-item label="密码:">-->
<!--            <span>{{ user_data.password }}</span>-->
<!--          </el-form-item>-->
          <el-form-item label="邮箱:">
            <span>{{ user_data.email }}</span>
          </el-form-item>
          <el-form-item label="电话:">
            <span>{{ user_data.phone }}</span>
          </el-form-item>
          <el-form-item label="类型:">
            <span>{{ type }}</span>
          </el-form-item>
          <el-form-item label="性别:">
            <span>{{ user_data.gender ? '男' : '女' }}</span>
          </el-form-item>
        </el-form>
      </el-card>
    </div>
  </div>
</template>

<script>

export default {
  name: "user_data",
  data() {
    return {
      user_data: localStorage.getItem("user_data")
        ? JSON.parse(localStorage.getItem("user_data"))
        : {},
      type: localStorage.getItem("user_type")
        ? JSON.parse(localStorage.getItem("user_type"))
        : "未知"
    };
  },
  created() {
    if (localStorage.getItem("user_data") == null) {
      this.findCurUser();
    }
  },
  methods: {
    findCurUser() {
      this.request.get("/users/data").then((res) => {
        if (res.status === 200) {
          localStorage.setItem("user_data", JSON.stringify(res.data));
        } else {
          this.$message.error(res.message);
        }
      });
    }
  }

};
</script>

