# This file can be dropped into a project to help locate the Pimoroni Pico libraries
# It will also set up the required include and module search paths.

if (DEFINED ENV{PIMORONI_PICO_FETCH_FROM_GIT} AND (NOT PIMORONI_PICO_FETCH_FROM_GIT))
    set(PIMORONI_PICO_FETCH_FROM_GIT $ENV{PIMORONI_PICO_FETCH_FROM_GIT})
    message("Using PIMORONI_PICO_FETCH_FROM_GIT from environment ('${PIMORONI_PICO_FETCH_FROM_GIT}')")
endif ()

if (DEFINED ENV{PIMORONI_PICO_FETCH_FROM_GIT_PATH} AND (NOT PIMORONI_PICO_FETCH_FROM_GIT_PATH))
    set(PIMORONI_PICO_FETCH_FROM_GIT_PATH $ENV{PIMORONI_PICO_FETCH_FROM_GIT_PATH})
    message("Using PIMORONI_PICO_FETCH_FROM_GIT_PATH from environment ('${PIMORONI_PICO_FETCH_FROM_GIT_PATH}')")
endif ()

