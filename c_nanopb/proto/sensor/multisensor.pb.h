/* Automatically generated nanopb header */
/* Generated by nanopb-0.4.0-dev */

#ifndef PB_SENSOR_SENSOR_MULTISENSOR_PB_H_INCLUDED
#define PB_SENSOR_SENSOR_MULTISENSOR_PB_H_INCLUDED
#include <pb.h>

#include "sensor/base.pb.h"

/* @@protoc_insertion_point(includes) */
#if PB_PROTO_HEADER_VERSION != 40
#error Regenerate this file with the current version of nanopb generator.
#endif

#ifdef __cplusplus
extern "C" {
#endif

/* Struct definitions */
typedef struct _sensor_MultiSensorStatus {
    uint32_t sequence;
    uint32_t crc;
    sensor_Temperature temperature;
    sensor_Humidity humidity;
    sensor_AmbientLight ambient_light;
    sensor_Battery battery;
/* @@protoc_insertion_point(struct:sensor_MultiSensorStatus) */
} sensor_MultiSensorStatus;


/* Initializer values for message structs */
#define sensor_MultiSensorStatus_init_default    {0, 0, sensor_Temperature_init_default, sensor_Humidity_init_default, sensor_AmbientLight_init_default, sensor_Battery_init_default}
#define sensor_MultiSensorStatus_init_zero       {0, 0, sensor_Temperature_init_zero, sensor_Humidity_init_zero, sensor_AmbientLight_init_zero, sensor_Battery_init_zero}

/* Field tags (for use in manual encoding/decoding) */
#define sensor_MultiSensorStatus_sequence_tag    1
#define sensor_MultiSensorStatus_crc_tag         2
#define sensor_MultiSensorStatus_temperature_tag 3
#define sensor_MultiSensorStatus_humidity_tag    4
#define sensor_MultiSensorStatus_ambient_light_tag 5
#define sensor_MultiSensorStatus_battery_tag     6

/* Struct field encoding specification for nanopb */
#define sensor_MultiSensorStatus_FIELDLIST(X, a) \
X(a, STATIC, SINGULAR, UINT32, sequence, 1) \
X(a, STATIC, SINGULAR, UINT32, crc, 2) \
X(a, STATIC, SINGULAR, MESSAGE, temperature, 3) \
X(a, STATIC, SINGULAR, MESSAGE, humidity, 4) \
X(a, STATIC, SINGULAR, MESSAGE, ambient_light, 5) \
X(a, STATIC, SINGULAR, MESSAGE, battery, 6)
#define sensor_MultiSensorStatus_CALLBACK NULL
#define sensor_MultiSensorStatus_DEFAULT NULL
#define sensor_MultiSensorStatus_temperature_MSGTYPE sensor_Temperature
#define sensor_MultiSensorStatus_humidity_MSGTYPE sensor_Humidity
#define sensor_MultiSensorStatus_ambient_light_MSGTYPE sensor_AmbientLight
#define sensor_MultiSensorStatus_battery_MSGTYPE sensor_Battery

extern const pb_msgdesc_t sensor_MultiSensorStatus_msg;

/* Defines for backwards compatibility with code written before nanopb-0.4.0 */
#define sensor_MultiSensorStatus_fields &sensor_MultiSensorStatus_msg

/* Maximum encoded size of messages (where known) */
#define sensor_MultiSensorStatus_size            53

#ifdef __cplusplus
} /* extern "C" */
#endif
/* @@protoc_insertion_point(eof) */

#endif
