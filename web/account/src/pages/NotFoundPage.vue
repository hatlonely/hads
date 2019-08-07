<style>
.notfound {
  font-size: 20em;
  color: grey;
}
</style>

<template>
  <v-app>
    <v-content>
      <v-container fluid fill-height>
        <v-layout column align-center justify-center>
          <v-flex md12>
            <span class="notfound">404</span>
          </v-flex>
          <v-flex md12>
            <span>Not Found</span>
          </v-flex>
        </v-layout>
      </v-container>
    </v-content>
  </v-app>
</template>


<script>
export default {
  name: "NotFound"
};
</script>
