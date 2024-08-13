<script setup>
import {reactive,ref} from 'vue'
import {Greet} from '../../wailsjs/go/main/App'

const res = reactive({
  results: []
})
const input = ref('')

function greet() {
  Greet(input.value).then(result => {
    console.log(result)
    res.results.push(result)
  })
}
function clear(){
  res.results = []
}

</script>

<template>
  <div>
    <el-input
        v-model="input"
        class="w-50 m-2"
        placeholder="Type something"
    />
    <el-table border :data="res.results" style="width: 100%;height: 100%">
      <el-table-column prop="address" label="ip地址"  ></el-table-column>
      <el-table-column prop="ip" label="ip" ></el-table-column>
    </el-table>
    <el-button type="primary" @click="greet">go通信</el-button>
    <el-button type="primary" @click="clear">clear</el-button>

  </div>
</template>

<style scoped>
.result {
  height: 20px;
  line-height: 20px;
  margin: 1.5rem auto;
}

.input-box .btn {
  width: 60px;
  height: 30px;
  line-height: 30px;
  border-radius: 3px;
  border: none;
  margin: 0 0 0 20px;
  padding: 0 8px;
  cursor: pointer;
}

.input-box .btn:hover {
  background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
  color: #333333;
}

.input-box .input {
  border: none;
  border-radius: 3px;
  outline: none;
  height: 30px;
  line-height: 30px;
  padding: 0 10px;
  background-color: rgba(240, 240, 240, 1);
  -webkit-font-smoothing: antialiased;
}

.input-box .input:hover {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}

.input-box .input:focus {
  border: none;
  background-color: rgba(255, 255, 255, 1);
}
</style>
