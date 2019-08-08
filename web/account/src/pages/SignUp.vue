<template>
  <v-app>
    <v-content>
      <v-layout align-center justify-center row fill-height text-center>
        <v-flex xs12>
          <v-card class="mx-auto" width="748" flat color="rgba(0,0,0,0)">
            <transition name="slide-x-transition" mode="out-in">
              <router-view></router-view>
            </transition>
            <v-footer color="rgba(0, 0, 0, 0)">
              <v-flex text-center xs12>
                <div class="caption text-right">
                  <a href="#">隐私</a>&nbsp;&nbsp;&nbsp;&nbsp;
                  <a href="#">条款</a>&nbsp;&nbsp;&nbsp;&nbsp;
                  <a href="#">帮助</a>
                </div>
              </v-flex>
            </v-footer>
          </v-card>
        </v-flex>
      </v-layout>
    </v-content>
  </v-app>
</template>

<script>
export default {};
</script>
