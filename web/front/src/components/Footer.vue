<template>
  <v-footer padless color="#417CA9">
    <v-row justify="center" no-gutters>
      <v-col class="py-4 text-center white--text" cols="12">duryun-blog</v-col>

      <div class="text-center white--text">
        <a class="text-center white--text" href="https://beian.miit.gov.cn/">{{icp_record}}</a>
      </div>
    </v-row>
  </v-footer>
</template>

<script>
export default {
  data() {
    return {
      icp_record: ''
    }
  },
  mounted() {
    this.$root.$on('msg', (msg) => {
      this.icp_record = msg
    })
  }
}
</script>

<style>
</style>