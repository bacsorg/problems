project(bunsan_bacs_problems)

include(${CMAKE_SOURCE_DIR}/system-config.cmake OPTIONAL)

cmake_minimum_required(VERSION 2.8)

include(BunsanCommon)

bunsan_use_bunsan(common pm)
bunsan_use_boost(system filesystem serialization unit_test_framework)
bunsan_use_protobuf()

# Common plugins
# TODO allow user to specify additional plugins
set(problem_types single)
foreach(problem_type ${problem_types})
    bunsan_use_bunsan(bacs_problem_${problem_type})
endforeach()

bunsan_add_shared_library(${PROJECT_NAME}
    src/stub.cpp
)

bunsan_install_headers()
bunsan_install_targets(${PROJECT_NAME})

bunsan_include_tests()
