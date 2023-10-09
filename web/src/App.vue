<script setup lang="ts"></script>

<template>
  <v-app id="inspire">
    <v-system-bar>
      <v-spacer></v-spacer>

      <v-icon>mdi-square</v-icon>

      <v-icon>mdi-circle</v-icon>

      <v-icon>mdi-triangle</v-icon>
    </v-system-bar>

    <v-navigation-drawer color="grey-lighten-3" rail expand-on-hover width="320">
      <v-list>
        <v-list-item
          prepend-avatar="https://blog.mahoo12138.cn/images/avatar.png"
          title="Mahoo12138"
          subtitle="mahoo12138@gmail.com"
        >
        </v-list-item>
      </v-list>
      <v-divider class="mx-3 my-5"></v-divider>

      <v-list density="compact" nav>
        <v-list-item
          prepend-icon="mdi-folder"
          title="My Files"
          value="myfiles"
        ></v-list-item>
        <v-list-item
          prepend-icon="mdi-account-multiple"
          title="Shared with me"
          value="shared"
        ></v-list-item>
        <v-list-item
          prepend-icon="mdi-star"
          title="Starred"
          value="starred"
        ></v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-navigation-drawer width="244">
      <v-sheet color="grey-lighten-1" height="128" width="100%"></v-sheet>

      <v-list nav>
        <v-list-item v-for="n in 5" :key="n" :value="n" link>
          <template v-slot:prepend>
            <v-icon icon="mdi-account"></v-icon>
          </template>

          <v-list-item-title v-text="`Person ${n}`"></v-list-item-title>
        </v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-app-bar class="px-3" color="grey-lighten-4" flat height="72">
      <v-spacer></v-spacer>

      <v-responsive max-width="156">
        <v-text-field
          bg-color="grey-lighten-1"
          density="compact"
          flat
          hide-details
          rounded="pill"
          variant="solo-filled"
        ></v-text-field>
      </v-responsive>
    </v-app-bar>

    <v-main><!--  --></v-main>

    <v-navigation-drawer location="right">
      <v-list>
        <v-list-item
          v-for="n in 12"
          :key="n"
          :title="`2022-${n}`"
          link
        ></v-list-item>
      </v-list>
    </v-navigation-drawer>

    <v-footer app height="72">
      <v-text-field
        bg-color="grey-lighten-1"
        class="overflow-hidden"
        density="compact"
        flat
        hide-details
        rounded="pill"
        variant="solo-filled"
      ></v-text-field>
    </v-footer>
  </v-app>
</template>

<style scoped></style>
