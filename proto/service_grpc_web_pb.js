/**
 * @fileoverview gRPC-Web generated client stub for main
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');

const proto = {};
proto.main = require('./service_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.main.HardwareMonitorClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.main.HardwareMonitorPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.main.EmptyRequest,
 *   !proto.main.HardwareStats>}
 */
const methodDescriptor_HardwareMonitor_Monitor = new grpc.web.MethodDescriptor(
  '/main.HardwareMonitor/Monitor',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.main.EmptyRequest,
  proto.main.HardwareStats,
  /**
   * @param {!proto.main.EmptyRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.main.HardwareStats.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.main.EmptyRequest,
 *   !proto.main.HardwareStats>}
 */
const methodInfo_HardwareMonitor_Monitor = new grpc.web.AbstractClientBase.MethodInfo(
  proto.main.HardwareStats,
  /**
   * @param {!proto.main.EmptyRequest} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.main.HardwareStats.deserializeBinary
);


/**
 * @param {!proto.main.EmptyRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.main.HardwareStats>}
 *     The XHR Node Readable Stream
 */
proto.main.HardwareMonitorClient.prototype.monitor =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/main.HardwareMonitor/Monitor',
      request,
      metadata || {},
      methodDescriptor_HardwareMonitor_Monitor);
};


/**
 * @param {!proto.main.EmptyRequest} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.main.HardwareStats>}
 *     The XHR Node Readable Stream
 */
proto.main.HardwareMonitorPromiseClient.prototype.monitor =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/main.HardwareMonitor/Monitor',
      request,
      metadata || {},
      methodDescriptor_HardwareMonitor_Monitor);
};


module.exports = proto.main;

