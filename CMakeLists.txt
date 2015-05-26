cmake_minimum_required(VERSION 3.0)

include(${CMAKE_SOURCE_DIR}/system-config.cmake OPTIONAL)

project(bunsan_bacs_problems)

if(BunsanCMakeBuiltin)
    BunsanCMakeLoad()
else()
    find_package(BunsanCMake CONFIG REQUIRED)
endif()

# Common plugins
# TODO allow user to specify additional plugins
set(problem_types
    single
)

bunsan_add_library(${PROJECT_NAME}
    src/stub.cpp
)
foreach(problem_type ${problem_types})
    set(package bunsan_bacs_problem_${problem_type})
    bunsan_use_bunsan_package(${PROJECT_NAME} ${package} ${package})
endforeach()

#bunsan_install_headers()
bunsan_install_targets(${PROJECT_NAME})
bunsan_install_project()

bunsan_include_tests()
