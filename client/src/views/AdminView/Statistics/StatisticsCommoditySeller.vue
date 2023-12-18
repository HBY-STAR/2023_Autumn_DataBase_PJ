<template>
  <div>
    <el-input-number
      v-model="commodity_id"
      style="width: 300px"
      :min="0"
      controls-position="right"
      placeholder="请输入可发布商品的序号"
    />
    <el-button @click="renderChart">
      <el-icon><Search /></el-icon>
    </el-button>
    <div style="float: right">
      <span style="color: #c45656">最高价商家名：{{this.chartData.labels[result.maxIndex]}}</span>
      <span style="color: #67c23a; margin-left: 40px;">最低价商家名：{{this.chartData.labels[result.minIndex]}}</span>
    </div>
  </div>
  <div style="margin-top: 20px">
    <canvas id="myChart" width="150" height="60"></canvas>
  </div>
</template>

<script>
import { Chart } from 'chart.js/auto';
import { Search } from "@element-plus/icons-vue";

export default {
  name: "statistics_price_seller",
  components: { Search },
  data() {
    return {
      chartData: {
        labels: [],
        datasets: [
          {
            label: '',
            backgroundColor: [
              'rgba(255, 99, 132, 0.2)',
              'rgba(255, 159, 64, 0.2)',
              'rgba(255, 205, 86, 0.2)',
              'rgba(75, 192, 192, 0.2)',
              'rgba(54, 162, 235, 0.2)',
              'rgba(153, 102, 255, 0.2)',
              'rgba(201, 203, 207, 0.2)'
            ],
            borderColor: [
              'rgb(255, 99, 132)',
              'rgb(255, 159, 64)',
              'rgb(255, 205, 86)',
              'rgb(75, 192, 192)',
              'rgb(54, 162, 235)',
              'rgb(153, 102, 255)',
              'rgb(201, 203, 207)'
            ],
            borderWidth: 1,
            data: [],
          },
        ],
      },
      commodity_id: null,
    };
  },
  mounted() {
    this.renderChart();
  },
  computed: {
    result() {
      return this.calculateMinMaxIndex(this.chartData.datasets[0].data);
    },
  },
  methods: {
    renderChart() {
      const ctx = document.getElementById('myChart').getContext('2d');
      // 销毁之前的图表实例
      if (this.chartInstance) {
        this.chartInstance.destroy();
      }

      if(this.commodity_id!=null){
        this.SearchStatisticsPriceSeller()
      }

      // 创建新的图表实例
      this.chartInstance = new Chart(ctx, {
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
    SearchStatisticsPriceSeller(){
      if(this.commodity_id==null){
        this.$message.error('请输入商品ID');
      }else {
        this.request.get("/item/commodity/"+this.commodity_id).then((res) => {
          if (res.status === 200) {
            if(res.data.length<1){
              this.$message.warning('数据为空');
            }else {
              this.chartData.labels=[];
              this.chartData.datasets=[
                {
                  label: '默认名称为 '+res.data[0].Commodity.default_name+" 的商品在不同商家的价格",
                  backgroundColor: [
                    'rgba(255, 99, 132, 0.2)',
                    'rgba(255, 159, 64, 0.2)',
                    'rgba(255, 205, 86, 0.2)',
                    'rgba(75, 192, 192, 0.2)',
                    'rgba(54, 162, 235, 0.2)',
                    'rgba(153, 102, 255, 0.2)',
                    'rgba(201, 203, 207, 0.2)'
                  ],
                  borderColor: [
                    'rgb(255, 99, 132)',
                    'rgb(255, 159, 64)',
                    'rgb(255, 205, 86)',
                    'rgb(75, 192, 192)',
                    'rgb(54, 162, 235)',
                    'rgb(153, 102, 255)',
                    'rgb(201, 203, 207)'
                  ],
                  borderWidth: 1,
                  data: [],
                },
              ];
              res.data.forEach(item => {
                this.chartData.labels.push(item.Seller.username);
              });
              res.data.forEach(item => {
                this.chartData.datasets[0].data.push(item.price);
              });
            }

          }
        });
      }
    },
    calculateMinMaxIndex(arr) {
      if (arr.length === 0) {
        return { maxIndex: null, minIndex: null };
      }

      let maxIndex = 0;
      let minIndex = 0;
      let maxValue = arr[0];
      let minValue = arr[0];

      for (let i = 1; i < arr.length; i++) {
        if (arr[i] > maxValue) {
          maxValue = arr[i];
          maxIndex = i;
        }

        if (arr[i] < minValue) {
          minValue = arr[i];
          minIndex = i;
        }
      }

      return {
        maxIndex: maxIndex,
        minIndex: minIndex,
      };
    },
  },
};
</script>
