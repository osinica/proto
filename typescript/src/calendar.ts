/* eslint-disable */
import * as _m0 from "protobufjs/minimal";

export const protobufPackage = "calendar";

export enum Status {
  SUCCESS = 0,
  ERROR = 1,
  UNRECOGNIZED = -1,
}

export function statusFromJSON(object: any): Status {
  switch (object) {
    case 0:
    case "SUCCESS":
      return Status.SUCCESS;
    case 1:
    case "ERROR":
      return Status.ERROR;
    case -1:
    case "UNRECOGNIZED":
    default:
      return Status.UNRECOGNIZED;
  }
}

export function statusToJSON(object: Status): string {
  switch (object) {
    case Status.SUCCESS:
      return "SUCCESS";
    case Status.ERROR:
      return "ERROR";
    case Status.UNRECOGNIZED:
    default:
      return "UNRECOGNIZED";
  }
}

export interface Event {
  id: string;
  name: string;
  start: number;
  end: number;
}

/** create */
export interface CreateEventRequest {
  name: string;
  start: number;
  ent: number;
}

export interface CreateEventResponse {
  status: Status;
  error: string;
  data: Event | undefined;
}

/** find one */
export interface FindOneRequest {
  id: string;
}

export interface FindOneResponse {
  status: Status;
  error: string;
  data: Event | undefined;
}

function createBaseEvent(): Event {
  return { id: "", name: "", start: 0, end: 0 };
}

export const Event = {
  encode(message: Event, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.start !== 0) {
      writer.uint32(24).uint32(message.start);
    }
    if (message.end !== 0) {
      writer.uint32(32).uint32(message.end);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Event {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseEvent();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.name = reader.string();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.start = reader.uint32();
          continue;
        case 4:
          if (tag !== 32) {
            break;
          }

          message.end = reader.uint32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): Event {
    return {
      id: isSet(object.id) ? String(object.id) : "",
      name: isSet(object.name) ? String(object.name) : "",
      start: isSet(object.start) ? Number(object.start) : 0,
      end: isSet(object.end) ? Number(object.end) : 0,
    };
  },

  toJSON(message: Event): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    message.name !== undefined && (obj.name = message.name);
    message.start !== undefined && (obj.start = Math.round(message.start));
    message.end !== undefined && (obj.end = Math.round(message.end));
    return obj;
  },

  create<I extends Exact<DeepPartial<Event>, I>>(base?: I): Event {
    return Event.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<Event>, I>>(object: I): Event {
    const message = createBaseEvent();
    message.id = object.id ?? "";
    message.name = object.name ?? "";
    message.start = object.start ?? 0;
    message.end = object.end ?? 0;
    return message;
  },
};

function createBaseCreateEventRequest(): CreateEventRequest {
  return { name: "", start: 0, ent: 0 };
}

export const CreateEventRequest = {
  encode(message: CreateEventRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.name !== "") {
      writer.uint32(10).string(message.name);
    }
    if (message.start !== 0) {
      writer.uint32(16).uint32(message.start);
    }
    if (message.ent !== 0) {
      writer.uint32(24).uint32(message.ent);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateEventRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateEventRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.name = reader.string();
          continue;
        case 2:
          if (tag !== 16) {
            break;
          }

          message.start = reader.uint32();
          continue;
        case 3:
          if (tag !== 24) {
            break;
          }

          message.ent = reader.uint32();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateEventRequest {
    return {
      name: isSet(object.name) ? String(object.name) : "",
      start: isSet(object.start) ? Number(object.start) : 0,
      ent: isSet(object.ent) ? Number(object.ent) : 0,
    };
  },

  toJSON(message: CreateEventRequest): unknown {
    const obj: any = {};
    message.name !== undefined && (obj.name = message.name);
    message.start !== undefined && (obj.start = Math.round(message.start));
    message.ent !== undefined && (obj.ent = Math.round(message.ent));
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateEventRequest>, I>>(base?: I): CreateEventRequest {
    return CreateEventRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateEventRequest>, I>>(object: I): CreateEventRequest {
    const message = createBaseCreateEventRequest();
    message.name = object.name ?? "";
    message.start = object.start ?? 0;
    message.ent = object.ent ?? 0;
    return message;
  },
};

function createBaseCreateEventResponse(): CreateEventResponse {
  return { status: 0, error: "", data: undefined };
}

export const CreateEventResponse = {
  encode(message: CreateEventResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.status !== 0) {
      writer.uint32(8).int32(message.status);
    }
    if (message.error !== "") {
      writer.uint32(18).string(message.error);
    }
    if (message.data !== undefined) {
      Event.encode(message.data, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): CreateEventResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseCreateEventResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.status = reader.int32() as any;
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.error = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.data = Event.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): CreateEventResponse {
    return {
      status: isSet(object.status) ? statusFromJSON(object.status) : 0,
      error: isSet(object.error) ? String(object.error) : "",
      data: isSet(object.data) ? Event.fromJSON(object.data) : undefined,
    };
  },

  toJSON(message: CreateEventResponse): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = statusToJSON(message.status));
    message.error !== undefined && (obj.error = message.error);
    message.data !== undefined && (obj.data = message.data ? Event.toJSON(message.data) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<CreateEventResponse>, I>>(base?: I): CreateEventResponse {
    return CreateEventResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<CreateEventResponse>, I>>(object: I): CreateEventResponse {
    const message = createBaseCreateEventResponse();
    message.status = object.status ?? 0;
    message.error = object.error ?? "";
    message.data = (object.data !== undefined && object.data !== null) ? Event.fromPartial(object.data) : undefined;
    return message;
  },
};

function createBaseFindOneRequest(): FindOneRequest {
  return { id: "" };
}

export const FindOneRequest = {
  encode(message: FindOneRequest, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.id !== "") {
      writer.uint32(10).string(message.id);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FindOneRequest {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFindOneRequest();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 10) {
            break;
          }

          message.id = reader.string();
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): FindOneRequest {
    return { id: isSet(object.id) ? String(object.id) : "" };
  },

  toJSON(message: FindOneRequest): unknown {
    const obj: any = {};
    message.id !== undefined && (obj.id = message.id);
    return obj;
  },

  create<I extends Exact<DeepPartial<FindOneRequest>, I>>(base?: I): FindOneRequest {
    return FindOneRequest.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<FindOneRequest>, I>>(object: I): FindOneRequest {
    const message = createBaseFindOneRequest();
    message.id = object.id ?? "";
    return message;
  },
};

function createBaseFindOneResponse(): FindOneResponse {
  return { status: 0, error: "", data: undefined };
}

export const FindOneResponse = {
  encode(message: FindOneResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.status !== 0) {
      writer.uint32(8).int32(message.status);
    }
    if (message.error !== "") {
      writer.uint32(18).string(message.error);
    }
    if (message.data !== undefined) {
      Event.encode(message.data, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): FindOneResponse {
    const reader = input instanceof _m0.Reader ? input : _m0.Reader.create(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseFindOneResponse();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          if (tag !== 8) {
            break;
          }

          message.status = reader.int32() as any;
          continue;
        case 2:
          if (tag !== 18) {
            break;
          }

          message.error = reader.string();
          continue;
        case 3:
          if (tag !== 26) {
            break;
          }

          message.data = Event.decode(reader, reader.uint32());
          continue;
      }
      if ((tag & 7) === 4 || tag === 0) {
        break;
      }
      reader.skipType(tag & 7);
    }
    return message;
  },

  fromJSON(object: any): FindOneResponse {
    return {
      status: isSet(object.status) ? statusFromJSON(object.status) : 0,
      error: isSet(object.error) ? String(object.error) : "",
      data: isSet(object.data) ? Event.fromJSON(object.data) : undefined,
    };
  },

  toJSON(message: FindOneResponse): unknown {
    const obj: any = {};
    message.status !== undefined && (obj.status = statusToJSON(message.status));
    message.error !== undefined && (obj.error = message.error);
    message.data !== undefined && (obj.data = message.data ? Event.toJSON(message.data) : undefined);
    return obj;
  },

  create<I extends Exact<DeepPartial<FindOneResponse>, I>>(base?: I): FindOneResponse {
    return FindOneResponse.fromPartial(base ?? {});
  },

  fromPartial<I extends Exact<DeepPartial<FindOneResponse>, I>>(object: I): FindOneResponse {
    const message = createBaseFindOneResponse();
    message.status = object.status ?? 0;
    message.error = object.error ?? "";
    message.data = (object.data !== undefined && object.data !== null) ? Event.fromPartial(object.data) : undefined;
    return message;
  },
};

export interface CalendarService {
  CreateEvent(request: CreateEventRequest): Promise<CreateEventResponse>;
  FindOne(request: FindOneRequest): Promise<FindOneResponse>;
}

export class CalendarServiceClientImpl implements CalendarService {
  private readonly rpc: Rpc;
  private readonly service: string;
  constructor(rpc: Rpc, opts?: { service?: string }) {
    this.service = opts?.service || "calendar.CalendarService";
    this.rpc = rpc;
    this.CreateEvent = this.CreateEvent.bind(this);
    this.FindOne = this.FindOne.bind(this);
  }
  CreateEvent(request: CreateEventRequest): Promise<CreateEventResponse> {
    const data = CreateEventRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "CreateEvent", data);
    return promise.then((data) => CreateEventResponse.decode(_m0.Reader.create(data)));
  }

  FindOne(request: FindOneRequest): Promise<FindOneResponse> {
    const data = FindOneRequest.encode(request).finish();
    const promise = this.rpc.request(this.service, "FindOne", data);
    return promise.then((data) => FindOneResponse.decode(_m0.Reader.create(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
