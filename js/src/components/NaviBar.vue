<template>
  <nav class="shadow-md">
    <div class="relative flex items-center justify-between h-12">
      <div
        class="
          flex-1 flex
          items-center
          justify-center
          sm:items-stretch sm:justify-start
        "
      >
        <div class="flex-shrink-0 flex items-center">
          <img
            class="block h-4 w-auto"
            src="../assets/konnect-logo-outline.svg"
            alt="Konnect"
          />
        </div>
      </div>
      <div
        class="
          absolute
          inset-y-0
          right-0
          flex
          items-center
          pr-2
          sm:static sm:inset-auto sm:ml-6 sm:pr-0
        "
      >
        <div class="flex gap-4">
        <KIcon size="20" icon="notificationBell" />
        <KIcon size="20" icon="help" />
        <button
          type="button"
          class="
            flex
            text-sm
            rounded-full
            focus:outline-none
            focus:ring-2
            focus:ring-offset-2
            focus:ring-offset-gray-800
            focus:ring-white
          "
          id="user-menu-button"
          aria-expanded="false"
          aria-haspopup="true"
        >
          <img class="h-4 w-4 rounded-full" src="" alt="" />
        </button>
        <span class="w-2"></span>
        </div>
      </div>
    </div>
  </nav>
</template>

<script>
import KIcon from "@kongponents/kicon";

export default {
  name: "NaviBar",
  components: { KIcon },
};
</script>
