package apk

import "github.com/lsongdev/apk-go/binary"

// Instrumentation is an application instrumentation code.
type Instrumentation struct {
	Name            binary.String `xml:"http://schemas.android.com/apk/res/android name,attr"`
	Target          binary.String `xml:"http://schemas.android.com/apk/res/android targetPackage,attr"`
	HandleProfiling binary.Bool   `xml:"http://schemas.android.com/apk/res/android handleProfiling,attr"`
	FunctionalTest  binary.Bool   `xml:"http://schemas.android.com/apk/res/android functionalTest,attr"`
}

// ActivityAction is an action of an activity.
type ActivityAction struct {
	Name binary.String `xml:"http://schemas.android.com/apk/res/android name,attr"`
}

// ActivityCategory is a category of an activity.
type ActivityCategory struct {
	Name binary.String `xml:"http://schemas.android.com/apk/res/android name,attr"`
}

// ActivityIntentFilter is an binary.Int32ent filter of an activity.
type ActivityIntentFilter struct {
	Actions    []ActivityAction   `xml:"action"`
	Categories []ActivityCategory `xml:"category"`
}

// AppActivity is an activity in an application.
type AppActivity struct {
	Theme             binary.String          `xml:"http://schemas.android.com/apk/res/android theme,attr"`
	Name              binary.String          `xml:"http://schemas.android.com/apk/res/android name,attr"`
	Label             binary.String          `xml:"http://schemas.android.com/apk/res/android label,attr"`
	ScreenOrientation binary.String          `xml:"http://schemas.android.com/apk/res/android screenOrientation,attr"`
	IntentFilters     []ActivityIntentFilter `xml:"intent-filter"`
}

// AppActivityAlias https://developer.android.com/guide/topics/manifest/activity-alias-element
type AppActivityAlias struct {
	Name           binary.String          `xml:"http://schemas.android.com/apk/res/android name,attr"`
	Label          binary.String          `xml:"http://schemas.android.com/apk/res/android label,attr"`
	TargetActivity binary.String          `xml:"http://schemas.android.com/apk/res/android targetActivity,attr"`
	IntentFilters  []ActivityIntentFilter `xml:"intent-filter"`
}

// MetaData is a metadata in an application.
type MetaData struct {
	Name  binary.String `xml:"http://schemas.android.com/apk/res/android name,attr"`
	Value binary.String `xml:"http://schemas.android.com/apk/res/android value,attr"`
}

// Application is an application in an APK.
type Application struct {
	AllowTaskReparenting  binary.Bool        `xml:"http://schemas.android.com/apk/res/android allowTaskReparenting,attr"`
	AllowBackup           binary.Bool        `xml:"http://schemas.android.com/apk/res/android allowBackup,attr"`
	BackupAgent           binary.String      `xml:"http://schemas.android.com/apk/res/android backupAgent,attr"`
	Debuggable            binary.Bool        `xml:"http://schemas.android.com/apk/res/android debuggable,attr"`
	Description           binary.String      `xml:"http://schemas.android.com/apk/res/android description,attr"`
	Enabled               binary.Bool        `xml:"http://schemas.android.com/apk/res/android enabled,attr"`
	HasCode               binary.Bool        `xml:"http://schemas.android.com/apk/res/android hasCode,attr"`
	HardwareAccelerated   binary.Bool        `xml:"http://schemas.android.com/apk/res/android hardwareAccelerated,attr"`
	Icon                  binary.String      `xml:"http://schemas.android.com/apk/res/android icon,attr"`
	KillAfterRestore      binary.Bool        `xml:"http://schemas.android.com/apk/res/android killAfterRestore,attr"`
	LargeHeap             binary.Bool        `xml:"http://schemas.android.com/apk/res/android largeHeap,attr"`
	Label                 binary.String      `xml:"http://schemas.android.com/apk/res/android label,attr"`
	Logo                  binary.String      `xml:"http://schemas.android.com/apk/res/android logo,attr"`
	ManageSpaceActivity   binary.String      `xml:"http://schemas.android.com/apk/res/android manageSpaceActivity,attr"`
	Name                  binary.String      `xml:"http://schemas.android.com/apk/res/android name,attr"`
	Permission            binary.String      `xml:"http://schemas.android.com/apk/res/android permission,attr"`
	Persistent            binary.Bool        `xml:"http://schemas.android.com/apk/res/android persistent,attr"`
	Process               binary.String      `xml:"http://schemas.android.com/apk/res/android process,attr"`
	RestoreAnyVersion     binary.Bool        `xml:"http://schemas.android.com/apk/res/android restoreAnyVersion,attr"`
	RequiredAccountType   binary.String      `xml:"http://schemas.android.com/apk/res/android requiredAccountType,attr"`
	RestrictedAccountType binary.String      `xml:"http://schemas.android.com/apk/res/android restrictedAccountType,attr"`
	SupportsRtl           binary.Bool        `xml:"http://schemas.android.com/apk/res/android supportsRtl,attr"`
	TaskAffinity          binary.String      `xml:"http://schemas.android.com/apk/res/android taskAffinity,attr"`
	TestOnly              binary.Bool        `xml:"http://schemas.android.com/apk/res/android testOnly,attr"`
	Theme                 binary.String      `xml:"http://schemas.android.com/apk/res/android theme,attr"`
	UIOptions             binary.String      `xml:"http://schemas.android.com/apk/res/android uiOptions,attr"`
	VMSafeMode            binary.Bool        `xml:"http://schemas.android.com/apk/res/android vmSafeMode,attr"`
	Activities            []AppActivity      `xml:"activity"`
	ActivityAliases       []AppActivityAlias `xml:"activity-alias"`
	MetaData              []MetaData         `xml:"meta-data"`
}

// UsesSDK is target SDK version.
type UsesSDK struct {
	Min    binary.Int32 `xml:"http://schemas.android.com/apk/res/android minSdkVersion,attr"`
	Target binary.Int32 `xml:"http://schemas.android.com/apk/res/android targetSdkVersion,attr"`
	Max    binary.Int32 `xml:"http://schemas.android.com/apk/res/android maxSdkVersion,attr"`
}

// UsesPermission is user grant the system permission.
type UsesPermission struct {
	Name binary.String `xml:"http://schemas.android.com/apk/res/android name,attr"`
	Max  binary.Int32  `xml:"http://schemas.android.com/apk/res/android maxSdkVersion,attr"`
}

// Manifest is a manifest of an APK.
type Manifest struct {
	Package                   binary.String    `xml:"package,attr"`
	CompileSDKVersion         binary.Int32     `xml:"http://schemas.android.com/apk/res/android compileSdkVersion,attr"`
	CompileSDKVersionCodename binary.String    `xml:"http://schemas.android.com/apk/res/android compileSdkVersionCodename,attr"`
	VersionCode               binary.Int32     `xml:"http://schemas.android.com/apk/res/android versionCode,attr"`
	VersionName               binary.String    `xml:"http://schemas.android.com/apk/res/android versionName,attr"`
	App                       Application      `xml:"application"`
	Instrument                Instrumentation  `xml:"instrumentation"`
	SDK                       UsesSDK          `xml:"uses-sdk"`
	UsesPermissions           []UsesPermission `xml:"uses-permission"`
}
