// Generated by the protocol buffer compiler. DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: test.proto

// Generated files should ignore deprecation warnings
@file:Suppress("DEPRECATION")
package li.xiam.uuidhelper.test;

@kotlin.jvm.JvmName("-initializechild")
public inline fun child(block: li.xiam.uuidhelper.test.ChildKt.Dsl.() -> kotlin.Unit): li.xiam.uuidhelper.test.Test.Child =
  li.xiam.uuidhelper.test.ChildKt.Dsl._create(li.xiam.uuidhelper.test.Test.Child.newBuilder()).apply { block() }._build()
/**
 * Protobuf type `uuidhelper.test.Child`
 */
public object ChildKt {
  @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
  @com.google.protobuf.kotlin.ProtoDslMarker
  public class Dsl private constructor(
    private val _builder: li.xiam.uuidhelper.test.Test.Child.Builder
  ) {
    public companion object {
      @kotlin.jvm.JvmSynthetic
    @kotlin.PublishedApi
      internal fun _create(builder: li.xiam.uuidhelper.test.Test.Child.Builder): Dsl = Dsl(builder)
    }

    @kotlin.jvm.JvmSynthetic
  @kotlin.PublishedApi
    internal fun _build(): li.xiam.uuidhelper.test.Test.Child = _builder.build()

    /**
     * `bytes child_uuid = 1;`
     */
    public var childUuid: com.google.protobuf.ByteString
      @JvmName("getChildUuid")
      get() = _builder.childUuid
      @JvmName("setChildUuid")
      set(value) {
        _builder.childUuid = value
      }
    /**
     * `bytes child_uuid = 1;`
     */
    public fun clearChildUuid() {
      _builder.clearChildUuid()
    }

    /**
     * An uninstantiable, behaviorless type to represent the field in
     * generics.
     */
    @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
    public class ChildUuidsProxy private constructor() : com.google.protobuf.kotlin.DslProxy()
    /**
     * `repeated bytes child_uuids = 2;`
     */
     public val childUuids: com.google.protobuf.kotlin.DslList<com.google.protobuf.ByteString, ChildUuidsProxy>
      @kotlin.jvm.JvmSynthetic
      get() = com.google.protobuf.kotlin.DslList(
        _builder.childUuidsList
      )
    /**
     * `repeated bytes child_uuids = 2;`
     * @param value The childUuids to add.
     */
    @kotlin.jvm.JvmSynthetic
    @kotlin.jvm.JvmName("addChildUuids")
    public fun com.google.protobuf.kotlin.DslList<com.google.protobuf.ByteString, ChildUuidsProxy>.add(value: com.google.protobuf.ByteString) {
      _builder.addChildUuids(value)
    }/**
     * `repeated bytes child_uuids = 2;`
     * @param value The childUuids to add.
     */
    @kotlin.jvm.JvmSynthetic
    @kotlin.jvm.JvmName("plusAssignChildUuids")
    @Suppress("NOTHING_TO_INLINE")
    public inline operator fun com.google.protobuf.kotlin.DslList<com.google.protobuf.ByteString, ChildUuidsProxy>.plusAssign(value: com.google.protobuf.ByteString) {
      add(value)
    }/**
     * `repeated bytes child_uuids = 2;`
     * @param values The childUuids to add.
     */
    @kotlin.jvm.JvmSynthetic
    @kotlin.jvm.JvmName("addAllChildUuids")
    public fun com.google.protobuf.kotlin.DslList<com.google.protobuf.ByteString, ChildUuidsProxy>.addAll(values: kotlin.collections.Iterable<com.google.protobuf.ByteString>) {
      _builder.addAllChildUuids(values)
    }/**
     * `repeated bytes child_uuids = 2;`
     * @param values The childUuids to add.
     */
    @kotlin.jvm.JvmSynthetic
    @kotlin.jvm.JvmName("plusAssignAllChildUuids")
    @Suppress("NOTHING_TO_INLINE")
    public inline operator fun com.google.protobuf.kotlin.DslList<com.google.protobuf.ByteString, ChildUuidsProxy>.plusAssign(values: kotlin.collections.Iterable<com.google.protobuf.ByteString>) {
      addAll(values)
    }/**
     * `repeated bytes child_uuids = 2;`
     * @param index The index to set the value at.
     * @param value The childUuids to set.
     */
    @kotlin.jvm.JvmSynthetic
    @kotlin.jvm.JvmName("setChildUuids")
    public operator fun com.google.protobuf.kotlin.DslList<com.google.protobuf.ByteString, ChildUuidsProxy>.set(index: kotlin.Int, value: com.google.protobuf.ByteString) {
      _builder.setChildUuids(index, value)
    }/**
     * `repeated bytes child_uuids = 2;`
     */
    @kotlin.jvm.JvmSynthetic
    @kotlin.jvm.JvmName("clearChildUuids")
    public fun com.google.protobuf.kotlin.DslList<com.google.protobuf.ByteString, ChildUuidsProxy>.clear() {
      _builder.clearChildUuids()
    }
    /**
     * An uninstantiable, behaviorless type to represent the field in
     * generics.
     */
    @kotlin.OptIn(com.google.protobuf.kotlin.OnlyForUseByGeneratedProtoCode::class)
    public class MapChildUuidProxy private constructor() : com.google.protobuf.kotlin.DslProxy()
    /**
     * `map<string, bytes> map_child_uuid = 3;`
     */
     public val mapChildUuid: com.google.protobuf.kotlin.DslMap<kotlin.String, com.google.protobuf.ByteString, MapChildUuidProxy>
      @kotlin.jvm.JvmSynthetic
      @JvmName("getMapChildUuidMap")
      get() = com.google.protobuf.kotlin.DslMap(
        _builder.mapChildUuidMap
      )
    /**
     * `map<string, bytes> map_child_uuid = 3;`
     */
    @JvmName("putMapChildUuid")
    public fun com.google.protobuf.kotlin.DslMap<kotlin.String, com.google.protobuf.ByteString, MapChildUuidProxy>
      .put(key: kotlin.String, value: com.google.protobuf.ByteString) {
         _builder.putMapChildUuid(key, value)
       }
    /**
     * `map<string, bytes> map_child_uuid = 3;`
     */
    @kotlin.jvm.JvmSynthetic
    @JvmName("setMapChildUuid")
    @Suppress("NOTHING_TO_INLINE")
    public inline operator fun com.google.protobuf.kotlin.DslMap<kotlin.String, com.google.protobuf.ByteString, MapChildUuidProxy>
      .set(key: kotlin.String, value: com.google.protobuf.ByteString) {
         put(key, value)
       }
    /**
     * `map<string, bytes> map_child_uuid = 3;`
     */
    @kotlin.jvm.JvmSynthetic
    @JvmName("removeMapChildUuid")
    public fun com.google.protobuf.kotlin.DslMap<kotlin.String, com.google.protobuf.ByteString, MapChildUuidProxy>
      .remove(key: kotlin.String) {
         _builder.removeMapChildUuid(key)
       }
    /**
     * `map<string, bytes> map_child_uuid = 3;`
     */
    @kotlin.jvm.JvmSynthetic
    @JvmName("putAllMapChildUuid")
    public fun com.google.protobuf.kotlin.DslMap<kotlin.String, com.google.protobuf.ByteString, MapChildUuidProxy>
      .putAll(map: kotlin.collections.Map<kotlin.String, com.google.protobuf.ByteString>) {
         _builder.putAllMapChildUuid(map)
       }
    /**
     * `map<string, bytes> map_child_uuid = 3;`
     */
    @kotlin.jvm.JvmSynthetic
    @JvmName("clearMapChildUuid")
    public fun com.google.protobuf.kotlin.DslMap<kotlin.String, com.google.protobuf.ByteString, MapChildUuidProxy>
      .clear() {
         _builder.clearMapChildUuid()
       }
  }
}
@kotlin.jvm.JvmSynthetic
public inline fun li.xiam.uuidhelper.test.Test.Child.copy(block: `li.xiam.uuidhelper.test`.ChildKt.Dsl.() -> kotlin.Unit): li.xiam.uuidhelper.test.Test.Child =
  `li.xiam.uuidhelper.test`.ChildKt.Dsl._create(this.toBuilder()).apply { block() }._build()

