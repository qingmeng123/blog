<template>
  <a-layout class="container">
    <Nav></Nav>

    <a-layout>
      <a-layout-header class="headerBtn">
        <Header></Header>
      </a-layout-header>
      <a-layout-content>
        <router-view :key="$route.fullPath"></router-view>
      </a-layout-content>
      <a-layout-footer class="footer">
        <Footer></Footer>
      </a-layout-footer>
    </a-layout>
  </a-layout>
</template>

<script>
import Nav from '../components/admin/Nav'
import Footer from '../components/admin/Footer'
import Header from '../components/admin/Header'

// function setRem() {
//   // 当前页面宽度相对于 750 宽的缩放比例，可根据自己需要修改
//   const scale = document.documentElement.clientWidth / 750;
//   document.documentElement.style.fontSize = scale + 'px';
// }
// setRem();
// window.onresize = setRem;
export default {
  components: { Nav, Footer, Header },
}
</script>

<style scoped>
.container {
  height: 100%;
}
.headerBtn {
  display: flex;
  justify-content: flex-end;
  align-items: center;
}
</style>
