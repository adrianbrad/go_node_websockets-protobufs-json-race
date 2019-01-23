/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!


goog.provide('proto.message.Message');
goog.provide('proto.message.TextNumberPair');

goog.require('jspb.BinaryReader');
goog.require('jspb.BinaryWriter');
goog.require('jspb.Message');


/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.message.Message = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.message.Message.repeatedFields_, null);
};
goog.inherits(proto.message.Message, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.message.Message.displayName = 'proto.message.Message';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.message.Message.repeatedFields_ = [3];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.message.Message.prototype.toObject = function(opt_includeInstance) {
  return proto.message.Message.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.message.Message} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.message.Message.toObject = function(includeInstance, msg) {
  var f, obj = {
    integer: jspb.Message.getFieldWithDefault(msg, 1, 0),
    floating: +jspb.Message.getFieldWithDefault(msg, 2, 0.0),
    pairsList: jspb.Message.toObjectList(msg.getPairsList(),
    proto.message.TextNumberPair.toObject, includeInstance)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.message.Message}
 */
proto.message.Message.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.message.Message;
  return proto.message.Message.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.message.Message} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.message.Message}
 */
proto.message.Message.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setInteger(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readDouble());
      msg.setFloating(value);
      break;
    case 3:
      var value = new proto.message.TextNumberPair;
      reader.readMessage(value,proto.message.TextNumberPair.deserializeBinaryFromReader);
      msg.addPairs(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.message.Message.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.message.Message.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.message.Message} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.message.Message.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getInteger();
  if (f !== 0) {
    writer.writeInt64(
      1,
      f
    );
  }
  f = message.getFloating();
  if (f !== 0.0) {
    writer.writeDouble(
      2,
      f
    );
  }
  f = message.getPairsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      3,
      f,
      proto.message.TextNumberPair.serializeBinaryToWriter
    );
  }
};


/**
 * optional int64 integer = 1;
 * @return {number}
 */
proto.message.Message.prototype.getInteger = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {number} value */
proto.message.Message.prototype.setInteger = function(value) {
  jspb.Message.setProto3IntField(this, 1, value);
};


/**
 * optional double floating = 2;
 * @return {number}
 */
proto.message.Message.prototype.getFloating = function() {
  return /** @type {number} */ (+jspb.Message.getFieldWithDefault(this, 2, 0.0));
};


/** @param {number} value */
proto.message.Message.prototype.setFloating = function(value) {
  jspb.Message.setProto3FloatField(this, 2, value);
};


/**
 * repeated TextNumberPair pairs = 3;
 * @return {!Array<!proto.message.TextNumberPair>}
 */
proto.message.Message.prototype.getPairsList = function() {
  return /** @type{!Array<!proto.message.TextNumberPair>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.message.TextNumberPair, 3));
};


/** @param {!Array<!proto.message.TextNumberPair>} value */
proto.message.Message.prototype.setPairsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 3, value);
};


/**
 * @param {!proto.message.TextNumberPair=} opt_value
 * @param {number=} opt_index
 * @return {!proto.message.TextNumberPair}
 */
proto.message.Message.prototype.addPairs = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 3, opt_value, proto.message.TextNumberPair, opt_index);
};


proto.message.Message.prototype.clearPairsList = function() {
  this.setPairsList([]);
};



/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.message.TextNumberPair = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.message.TextNumberPair, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.message.TextNumberPair.displayName = 'proto.message.TextNumberPair';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.message.TextNumberPair.prototype.toObject = function(opt_includeInstance) {
  return proto.message.TextNumberPair.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.message.TextNumberPair} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.message.TextNumberPair.toObject = function(includeInstance, msg) {
  var f, obj = {
    text: jspb.Message.getFieldWithDefault(msg, 1, ""),
    number: jspb.Message.getFieldWithDefault(msg, 2, 0)
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.message.TextNumberPair}
 */
proto.message.TextNumberPair.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.message.TextNumberPair;
  return proto.message.TextNumberPair.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.message.TextNumberPair} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.message.TextNumberPair}
 */
proto.message.TextNumberPair.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setText(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setNumber(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.message.TextNumberPair.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.message.TextNumberPair.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.message.TextNumberPair} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.message.TextNumberPair.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getText();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getNumber();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
};


/**
 * optional string text = 1;
 * @return {string}
 */
proto.message.TextNumberPair.prototype.getText = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.message.TextNumberPair.prototype.setText = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional int64 number = 2;
 * @return {number}
 */
proto.message.TextNumberPair.prototype.getNumber = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.message.TextNumberPair.prototype.setNumber = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};

