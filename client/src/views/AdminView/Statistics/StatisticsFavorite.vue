<template>
  <div>
    <el-select v-model="search_info.all" placeholder="是否指定用户种类" style="width: 180px;">
      <el-option label="默认所有用户" value= 'true' />
      <el-option label="指定用户种类" value= 'false' />
    </el-select>
    <el-select v-model="search_info.gender" placeholder="性别"  style="width: 100px">
      <el-option label="男" value= 'true' />
      <el-option label="女" value = 'false' />
    </el-select>
    <el-input-number
      v-model="search_info.age_start"
      style="width: 200px"
      :min="0"
      :max="150"
      controls-position="right"
      placeholder="请输入起始年龄"
    />
    <el-input-number
      v-model="search_info.age_end"
      style="width: 200px"
      :min="0"
      :max="150"
      controls-position="right"
      placeholder="请输入结束年龄"
    />
    <el-button @click="SearchStatusFavorite">
      <el-icon><Search /></el-icon>
    </el-button>
  </div>
  <div style="margin-top: 20px">
    <canvas id="myChart" width="200" height="80"></canvas>
  </div>
</template>

<script>
import { Chart } from 'chart.js/auto';
import { Search } from "@element-plus/icons-vue";

export default {
  name: "statistics_favorite",
  components: { Search },
  data() {
    return {
      chartData: {
        labels: [],
        datasets: [
          {
            label: '',
            backgroundColor: 'rgba(30,98,224,0.61)',
            borderColor: 'rgba(43,250,11,0.49)',
            borderWidth: 1,
            data: [],
          },
        ],
      },
      search_info:{
        all:null,
        gender: null,
        age_start: null,
        age_end: null,
      },
    };
  },
  created() {
    this.search_info.all='true'
    this.SearchStatusFavorite()
  },
  mounted() {
    this.renderChart();
  },
  methods: {
    renderChart() {
      const ctx = document.getElementById('myChart').getContext('2d');
      new Chart(ctx, {
        type: 'bar',
        data: this.chartData,
        options: {
          scales: {
            y: {
              beginAtZero: true,
            },
          },
        },
      });
    },
    SearchStatusFavorite(){
      this.search_info.all = this.search_info.all === 'true'
      if(this.search_info.all === true){
        this.search_info.gender=null
        this.search_info.age_start=null
        this.search_info.age_end=null
        // request
        this.request.post("/favorite/statistics",this.search_info).then((res) => {
          if (res.status === 200) {
            this.chartData.labels = res.data.CommodityItem.item_name
            this.chartData.datasets=[
              {
                label: '所有用户收藏前10的商品',
                backgroundColor: 'rgba(30,98,224,0.61)',
                borderColor: 'rgba(43,250,11,0.49)',
                borderWidth: 1,
                data: res.data.count
              },
            ]
          }
        }).catch(error => {
          this.$message.error(error.response.data.message);
        });
      }else {
        if(this.search_info.gender==null||this.search_info.age_start==null||this.search_info.age_end==null){
          this.$message.error('若指定用户范围，请指定完整的性别和年龄范围')
        }else{
          this.search_info.all = this.search_info.all === 'true'
          this.search_info.age_start=parseInt(this.search_info.age_start)
          this.search_info.age_end=parseInt(this.search_info.age_end)
          // request
          this.request.post("/favorite/statistics",this.search_info).then((res) => {
            if (res.status === 200) {
              this.chartData.labels = res.data.CommodityItem.item_name
              this.chartData.datasets=[
                {
                  label: '指定用户收藏前10的商品',
                  backgroundColor: 'rgba(30,98,224,0.61)',
                  borderColor: 'rgba(43,250,11,0.49)',
                  borderWidth: 1,
                  data: res.data.count
                },
              ]
            }
          }).catch(error => {
            this.$message.error(error.response.data.message);
          });
        }
      }
    },
  },
};
</script>
