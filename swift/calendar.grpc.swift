//
// DO NOT EDIT.
// swift-format-ignore-file
//
// Generated by the protocol buffer compiler.
// Source: proto/calendar.proto
//

//
// Copyright 2018, gRPC Authors All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
import GRPC
import NIO
import NIOConcurrencyHelpers
import SwiftProtobuf


/// Usage: instantiate `Calendar_CalendarServiceClient`, then call methods of this protocol to make API calls.
internal protocol Calendar_CalendarServiceClientProtocol: GRPCClient {
  var serviceName: String { get }
  var interceptors: Calendar_CalendarServiceClientInterceptorFactoryProtocol? { get }

  func createEvent(
    _ request: Calendar_CreateEventRequest,
    callOptions: CallOptions?
  ) -> UnaryCall<Calendar_CreateEventRequest, Calendar_CreateEventResponse>

  func findOne(
    _ request: Calendar_FindOneRequest,
    callOptions: CallOptions?
  ) -> UnaryCall<Calendar_FindOneRequest, Calendar_FindOneResponse>
}

extension Calendar_CalendarServiceClientProtocol {
  internal var serviceName: String {
    return "calendar.CalendarService"
  }

  /// Unary call to CreateEvent
  ///
  /// - Parameters:
  ///   - request: Request to send to CreateEvent.
  ///   - callOptions: Call options.
  /// - Returns: A `UnaryCall` with futures for the metadata, status and response.
  internal func createEvent(
    _ request: Calendar_CreateEventRequest,
    callOptions: CallOptions? = nil
  ) -> UnaryCall<Calendar_CreateEventRequest, Calendar_CreateEventResponse> {
    return self.makeUnaryCall(
      path: Calendar_CalendarServiceClientMetadata.Methods.createEvent.path,
      request: request,
      callOptions: callOptions ?? self.defaultCallOptions,
      interceptors: self.interceptors?.makeCreateEventInterceptors() ?? []
    )
  }

  /// Unary call to FindOne
  ///
  /// - Parameters:
  ///   - request: Request to send to FindOne.
  ///   - callOptions: Call options.
  /// - Returns: A `UnaryCall` with futures for the metadata, status and response.
  internal func findOne(
    _ request: Calendar_FindOneRequest,
    callOptions: CallOptions? = nil
  ) -> UnaryCall<Calendar_FindOneRequest, Calendar_FindOneResponse> {
    return self.makeUnaryCall(
      path: Calendar_CalendarServiceClientMetadata.Methods.findOne.path,
      request: request,
      callOptions: callOptions ?? self.defaultCallOptions,
      interceptors: self.interceptors?.makeFindOneInterceptors() ?? []
    )
  }
}

@available(*, deprecated)
extension Calendar_CalendarServiceClient: @unchecked Sendable {}

@available(*, deprecated, renamed: "Calendar_CalendarServiceNIOClient")
internal final class Calendar_CalendarServiceClient: Calendar_CalendarServiceClientProtocol {
  private let lock = Lock()
  private var _defaultCallOptions: CallOptions
  private var _interceptors: Calendar_CalendarServiceClientInterceptorFactoryProtocol?
  internal let channel: GRPCChannel
  internal var defaultCallOptions: CallOptions {
    get { self.lock.withLock { return self._defaultCallOptions } }
    set { self.lock.withLockVoid { self._defaultCallOptions = newValue } }
  }
  internal var interceptors: Calendar_CalendarServiceClientInterceptorFactoryProtocol? {
    get { self.lock.withLock { return self._interceptors } }
    set { self.lock.withLockVoid { self._interceptors = newValue } }
  }

  /// Creates a client for the calendar.CalendarService service.
  ///
  /// - Parameters:
  ///   - channel: `GRPCChannel` to the service host.
  ///   - defaultCallOptions: Options to use for each service call if the user doesn't provide them.
  ///   - interceptors: A factory providing interceptors for each RPC.
  internal init(
    channel: GRPCChannel,
    defaultCallOptions: CallOptions = CallOptions(),
    interceptors: Calendar_CalendarServiceClientInterceptorFactoryProtocol? = nil
  ) {
    self.channel = channel
    self._defaultCallOptions = defaultCallOptions
    self._interceptors = interceptors
  }
}

internal struct Calendar_CalendarServiceNIOClient: Calendar_CalendarServiceClientProtocol {
  internal var channel: GRPCChannel
  internal var defaultCallOptions: CallOptions
  internal var interceptors: Calendar_CalendarServiceClientInterceptorFactoryProtocol?

  /// Creates a client for the calendar.CalendarService service.
  ///
  /// - Parameters:
  ///   - channel: `GRPCChannel` to the service host.
  ///   - defaultCallOptions: Options to use for each service call if the user doesn't provide them.
  ///   - interceptors: A factory providing interceptors for each RPC.
  internal init(
    channel: GRPCChannel,
    defaultCallOptions: CallOptions = CallOptions(),
    interceptors: Calendar_CalendarServiceClientInterceptorFactoryProtocol? = nil
  ) {
    self.channel = channel
    self.defaultCallOptions = defaultCallOptions
    self.interceptors = interceptors
  }
}

@available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
internal protocol Calendar_CalendarServiceAsyncClientProtocol: GRPCClient {
  static var serviceDescriptor: GRPCServiceDescriptor { get }
  var interceptors: Calendar_CalendarServiceClientInterceptorFactoryProtocol? { get }

  func makeCreateEventCall(
    _ request: Calendar_CreateEventRequest,
    callOptions: CallOptions?
  ) -> GRPCAsyncUnaryCall<Calendar_CreateEventRequest, Calendar_CreateEventResponse>

  func makeFindOneCall(
    _ request: Calendar_FindOneRequest,
    callOptions: CallOptions?
  ) -> GRPCAsyncUnaryCall<Calendar_FindOneRequest, Calendar_FindOneResponse>
}

@available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
extension Calendar_CalendarServiceAsyncClientProtocol {
  internal static var serviceDescriptor: GRPCServiceDescriptor {
    return Calendar_CalendarServiceClientMetadata.serviceDescriptor
  }

  internal var interceptors: Calendar_CalendarServiceClientInterceptorFactoryProtocol? {
    return nil
  }

  internal func makeCreateEventCall(
    _ request: Calendar_CreateEventRequest,
    callOptions: CallOptions? = nil
  ) -> GRPCAsyncUnaryCall<Calendar_CreateEventRequest, Calendar_CreateEventResponse> {
    return self.makeAsyncUnaryCall(
      path: Calendar_CalendarServiceClientMetadata.Methods.createEvent.path,
      request: request,
      callOptions: callOptions ?? self.defaultCallOptions,
      interceptors: self.interceptors?.makeCreateEventInterceptors() ?? []
    )
  }

  internal func makeFindOneCall(
    _ request: Calendar_FindOneRequest,
    callOptions: CallOptions? = nil
  ) -> GRPCAsyncUnaryCall<Calendar_FindOneRequest, Calendar_FindOneResponse> {
    return self.makeAsyncUnaryCall(
      path: Calendar_CalendarServiceClientMetadata.Methods.findOne.path,
      request: request,
      callOptions: callOptions ?? self.defaultCallOptions,
      interceptors: self.interceptors?.makeFindOneInterceptors() ?? []
    )
  }
}

@available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
extension Calendar_CalendarServiceAsyncClientProtocol {
  internal func createEvent(
    _ request: Calendar_CreateEventRequest,
    callOptions: CallOptions? = nil
  ) async throws -> Calendar_CreateEventResponse {
    return try await self.performAsyncUnaryCall(
      path: Calendar_CalendarServiceClientMetadata.Methods.createEvent.path,
      request: request,
      callOptions: callOptions ?? self.defaultCallOptions,
      interceptors: self.interceptors?.makeCreateEventInterceptors() ?? []
    )
  }

  internal func findOne(
    _ request: Calendar_FindOneRequest,
    callOptions: CallOptions? = nil
  ) async throws -> Calendar_FindOneResponse {
    return try await self.performAsyncUnaryCall(
      path: Calendar_CalendarServiceClientMetadata.Methods.findOne.path,
      request: request,
      callOptions: callOptions ?? self.defaultCallOptions,
      interceptors: self.interceptors?.makeFindOneInterceptors() ?? []
    )
  }
}

@available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
internal struct Calendar_CalendarServiceAsyncClient: Calendar_CalendarServiceAsyncClientProtocol {
  internal var channel: GRPCChannel
  internal var defaultCallOptions: CallOptions
  internal var interceptors: Calendar_CalendarServiceClientInterceptorFactoryProtocol?

  internal init(
    channel: GRPCChannel,
    defaultCallOptions: CallOptions = CallOptions(),
    interceptors: Calendar_CalendarServiceClientInterceptorFactoryProtocol? = nil
  ) {
    self.channel = channel
    self.defaultCallOptions = defaultCallOptions
    self.interceptors = interceptors
  }
}

internal protocol Calendar_CalendarServiceClientInterceptorFactoryProtocol: Sendable {

  /// - Returns: Interceptors to use when invoking 'createEvent'.
  func makeCreateEventInterceptors() -> [ClientInterceptor<Calendar_CreateEventRequest, Calendar_CreateEventResponse>]

  /// - Returns: Interceptors to use when invoking 'findOne'.
  func makeFindOneInterceptors() -> [ClientInterceptor<Calendar_FindOneRequest, Calendar_FindOneResponse>]
}

internal enum Calendar_CalendarServiceClientMetadata {
  internal static let serviceDescriptor = GRPCServiceDescriptor(
    name: "CalendarService",
    fullName: "calendar.CalendarService",
    methods: [
      Calendar_CalendarServiceClientMetadata.Methods.createEvent,
      Calendar_CalendarServiceClientMetadata.Methods.findOne,
    ]
  )

  internal enum Methods {
    internal static let createEvent = GRPCMethodDescriptor(
      name: "CreateEvent",
      path: "/calendar.CalendarService/CreateEvent",
      type: GRPCCallType.unary
    )

    internal static let findOne = GRPCMethodDescriptor(
      name: "FindOne",
      path: "/calendar.CalendarService/FindOne",
      type: GRPCCallType.unary
    )
  }
}

/// To build a server, implement a class that conforms to this protocol.
internal protocol Calendar_CalendarServiceProvider: CallHandlerProvider {
  var interceptors: Calendar_CalendarServiceServerInterceptorFactoryProtocol? { get }

  func createEvent(request: Calendar_CreateEventRequest, context: StatusOnlyCallContext) -> EventLoopFuture<Calendar_CreateEventResponse>

  func findOne(request: Calendar_FindOneRequest, context: StatusOnlyCallContext) -> EventLoopFuture<Calendar_FindOneResponse>
}

extension Calendar_CalendarServiceProvider {
  internal var serviceName: Substring {
    return Calendar_CalendarServiceServerMetadata.serviceDescriptor.fullName[...]
  }

  /// Determines, calls and returns the appropriate request handler, depending on the request's method.
  /// Returns nil for methods not handled by this service.
  internal func handle(
    method name: Substring,
    context: CallHandlerContext
  ) -> GRPCServerHandlerProtocol? {
    switch name {
    case "CreateEvent":
      return UnaryServerHandler(
        context: context,
        requestDeserializer: ProtobufDeserializer<Calendar_CreateEventRequest>(),
        responseSerializer: ProtobufSerializer<Calendar_CreateEventResponse>(),
        interceptors: self.interceptors?.makeCreateEventInterceptors() ?? [],
        userFunction: self.createEvent(request:context:)
      )

    case "FindOne":
      return UnaryServerHandler(
        context: context,
        requestDeserializer: ProtobufDeserializer<Calendar_FindOneRequest>(),
        responseSerializer: ProtobufSerializer<Calendar_FindOneResponse>(),
        interceptors: self.interceptors?.makeFindOneInterceptors() ?? [],
        userFunction: self.findOne(request:context:)
      )

    default:
      return nil
    }
  }
}

/// To implement a server, implement an object which conforms to this protocol.
@available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
internal protocol Calendar_CalendarServiceAsyncProvider: CallHandlerProvider, Sendable {
  static var serviceDescriptor: GRPCServiceDescriptor { get }
  var interceptors: Calendar_CalendarServiceServerInterceptorFactoryProtocol? { get }

  func createEvent(
    request: Calendar_CreateEventRequest,
    context: GRPCAsyncServerCallContext
  ) async throws -> Calendar_CreateEventResponse

  func findOne(
    request: Calendar_FindOneRequest,
    context: GRPCAsyncServerCallContext
  ) async throws -> Calendar_FindOneResponse
}

@available(macOS 10.15, iOS 13, tvOS 13, watchOS 6, *)
extension Calendar_CalendarServiceAsyncProvider {
  internal static var serviceDescriptor: GRPCServiceDescriptor {
    return Calendar_CalendarServiceServerMetadata.serviceDescriptor
  }

  internal var serviceName: Substring {
    return Calendar_CalendarServiceServerMetadata.serviceDescriptor.fullName[...]
  }

  internal var interceptors: Calendar_CalendarServiceServerInterceptorFactoryProtocol? {
    return nil
  }

  internal func handle(
    method name: Substring,
    context: CallHandlerContext
  ) -> GRPCServerHandlerProtocol? {
    switch name {
    case "CreateEvent":
      return GRPCAsyncServerHandler(
        context: context,
        requestDeserializer: ProtobufDeserializer<Calendar_CreateEventRequest>(),
        responseSerializer: ProtobufSerializer<Calendar_CreateEventResponse>(),
        interceptors: self.interceptors?.makeCreateEventInterceptors() ?? [],
        wrapping: { try await self.createEvent(request: $0, context: $1) }
      )

    case "FindOne":
      return GRPCAsyncServerHandler(
        context: context,
        requestDeserializer: ProtobufDeserializer<Calendar_FindOneRequest>(),
        responseSerializer: ProtobufSerializer<Calendar_FindOneResponse>(),
        interceptors: self.interceptors?.makeFindOneInterceptors() ?? [],
        wrapping: { try await self.findOne(request: $0, context: $1) }
      )

    default:
      return nil
    }
  }
}

internal protocol Calendar_CalendarServiceServerInterceptorFactoryProtocol: Sendable {

  /// - Returns: Interceptors to use when handling 'createEvent'.
  ///   Defaults to calling `self.makeInterceptors()`.
  func makeCreateEventInterceptors() -> [ServerInterceptor<Calendar_CreateEventRequest, Calendar_CreateEventResponse>]

  /// - Returns: Interceptors to use when handling 'findOne'.
  ///   Defaults to calling `self.makeInterceptors()`.
  func makeFindOneInterceptors() -> [ServerInterceptor<Calendar_FindOneRequest, Calendar_FindOneResponse>]
}

internal enum Calendar_CalendarServiceServerMetadata {
  internal static let serviceDescriptor = GRPCServiceDescriptor(
    name: "CalendarService",
    fullName: "calendar.CalendarService",
    methods: [
      Calendar_CalendarServiceServerMetadata.Methods.createEvent,
      Calendar_CalendarServiceServerMetadata.Methods.findOne,
    ]
  )

  internal enum Methods {
    internal static let createEvent = GRPCMethodDescriptor(
      name: "CreateEvent",
      path: "/calendar.CalendarService/CreateEvent",
      type: GRPCCallType.unary
    )

    internal static let findOne = GRPCMethodDescriptor(
      name: "FindOne",
      path: "/calendar.CalendarService/FindOne",
      type: GRPCCallType.unary
    )
  }
}
