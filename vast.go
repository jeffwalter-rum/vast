// Package vast implements IAB VAST 4.2 https://iabtechlab.com/wp-content/uploads/2019/06/VAST_4.2_final_june26.pdf
package vast

import "encoding/xml"

// VAST is the root <VAST> tag
type VAST struct {
	// The version of the VAST spec (should be either "2.0" or "3.0")
	Version string `xml:"version,attr" json:",omitempty"`
	// XML namespace. Most likely 'http://www.iab.com/VAST'
	XMLNS string `xml:"xmlns,attr,omitempty" json:"xmlns,omitempty"`
	// One or more Ad elements. Advertisers and video content publishers may
	// associate an <Ad> element with a line item video ad defined in contract
	// documentation, usually an insertion order. These line item ads typically
	// specify the creative to display, price, delivery schedule, targeting,
	// and so on.
	Ads []Ad `xml:"Ad,omitempty" json:"Ad,omitempty"`
	// Contains a URI to a tracking resource that the video player should request
	// upon receiving a “no ad” response
	Errors []CDATAString `xml:"Error,omitempty" json:",omitempty"`

	Mute bool `xml:"mute,attr,omitempty" json:",omitempty"`
}

// Ad represent an <Ad> child tag in a VAST document
//
// Each <Ad> contains a single <InLine> element or <Wrapper> element (but never both).
type Ad struct {
	// InLine
	InLine *InLine `xml:",omitempty" json:",omitempty"`
	// Wrapper
	Wrapper *Wrapper `xml:",omitempty" json:",omitempty"`
	// ID is an ad server-defined identifier string for the ad
	ID string `xml:"id,attr,omitempty" json:",omitempty"`
	// AdType is an optional string that identifies the type of ad. This allows VAST to support audio ad scenarios.
	// Possible values – video, audio, hybrid.
	// Default value – video (assumed to be video if attribute is not present)
	AdType string `xml:"adType,attr,omitempty" json:",omitempty"`
	// Sequence is a number greater than zero (0) that identifies the sequence in which
	// an ad should play; all <Ad> elements with sequence values are part of
	// a pod and are intended to be played in sequence
	Sequence int `xml:"sequence,attr,omitempty" json:",omitempty"`
}

// CDATAString
// Written as character data wrapped in one or more <![CDATA[ ... ]]> tags, not as an XML element.
type CDATAString struct {
	CDATA string `xml:",cdata" json:"Data"`
}

// PlainString
// Written as plain character data, not as an XML element.
type PlainString struct {
	CDATA string `xml:",chardata" json:"Data"`
}

// InLine is a vast <InLine> ad element containing actual ad definition
// The last ad server in the ad supply chain serves an <InLine> element.
// Within the nested elements of an <InLine> element are all the files and
// URIs necessary to display the ad.
type InLine struct {
	// The name of the ad server that returned the ad
	AdSystem AdSystem
	// The common name of the ad
	AdTitle string
	// One or more URIs that directs the video player to a tracking resource file that the
	// video player should request when the first frame of the ad is displayed
	Impressions []Impression `xml:"Impression"`
	// Any ad server that returns a VAST containing an <InLine> ad must generate a pseudo-unique identifier
	// that is appropriate for all involved parties to track the lifecycle of that ad.
	// Example: ServerName-47ed3bac-1768-4b9a-9d0e-0b92422ab066
	AdServingId string `xml:",omitempty" json:",omitempty"`
	// Used in creative separation and for compliance in certain programs, a category field is
	// needed to categorize the ad’s content. Several category lists exist, some for describing site
	// content and some for describing ad content. Some lists are used interchangeably for both
	// site content and ad content. For example, the category list used to comply with the IAB
	// Quality Assurance Guidelines (QAG) describes site content, but is sometimes used to
	// describe ad content.
	// The VAST category field should only use AD CONTENT description categories.
	Category []Category `xml:",omitempty" json:",omitempty"`
	// A string value that provides a longer description of the ad.
	Description string `xml:",omitempty" json:",omitempty"`
	// The name of the advertiser as defined by the ad serving party.
	// This element can be used to prevent displaying ads with advertiser
	// competitors. Ad serving parties and publishers should identify how
	// to interpret values provided within this element. As with any optional
	// elements, the video player is not required to support it.
	Advertiser *Advertiser `xml:",omitempty" json:",omitempty"`
	// Provides a value that represents a price that can be used by real-time bidding
	// (RTB) systems. VAST is not designed to handle RTB since other methods exist,
	// but this element is offered for custom solutions if needed.
	Pricing *Pricing `xml:",omitempty" json:",omitempty"`
	// A URI to a survey vendor that could be the survey, a tracking pixel,
	// or anything to do with the survey. Multiple survey elements can be provided.
	// A type attribute is available to specify the MIME type being served.
	// For example, the attribute might be set to type=”text/javascript”.
	// Surveys can be dynamically inserted into the VAST response as long as
	// cross-domain issues are avoided.
	Survey *Survey `xml:",omitempty" json:",omitempty"`
	// A URI representing an error-tracking pixel; this element can occur multiple
	// times.
	Errors []CDATAString `xml:"Error,omitempty" json:"Error,omitempty"`
	// XML node for custom extensions, as defined by the ad server. When used, a
	// custom element should be nested under <Extensions> to help separate custom
	// XML elements from VAST elements. The following example includes a custom
	// xml element within the Extensions element.
	Extensions *[]Extension `xml:"Extensions>Extension,omitempty" json:",omitempty"`
	// The ad server may provide URIs for tracking publisher-determined viewability,
	// for both the InLine ad and any Wrappers, using the <ViewableImpression> element.
	// Tracking URIs may be provided in three containers: <Viewable>, <NotViewable>, and <ViewUndetermined>.
	ViewableImpression *ViewableImpression `xml:",omitempty" json:",omitempty"`
	// The <AdVerifications> element contains one or more <Verification> elements,
	// which list the resources and metadata required to execute third-party measurement code in order to verify creative playback.
	// The <AdVerifications> element is used to contain one or more <Verification> elements,
	// which are used to initiate a controlled container where code can be executed for collecting data to verify ad playback details.
	AdVerifications []Verification `xml:",omitempty" json:",omitempty"`
	// Provides a value that represents a price that can be used by real-time bidding
	// (RTB) systems. VAST is not designed to handle RTB since other methods exist,
	// but this element is offered for custom solutions if needed.
	// The container for one or more <Creative> elements
	Creatives []Creative `xml:"Creatives>Creative"`
	// The number of seconds in which the ad is valid for execution.
	// In cases where the ad is requested ahead of time, this timing indicates how many seconds after the request that the ad expires and cannot be played.
	// This element is useful for preventing an ad from playing after a timeout has occurred.
	Expires *int `xml:",omitempty" json:",omitempty"`
}

// Impression is a URI that directs the video player to a tracking resource file that
// the video player should request when the first frame of the ad is displayed
type Impression struct {
	ID  string `xml:"id,attr,omitempty" json:",omitempty"`
	URI string `xml:",cdata"`
}

// Pricing provides a value that represents a price that can be used by real-time
// bidding (RTB) systems. VAST is not designed to handle RTB since other methods
// exist,  but this element is offered for custom solutions if needed.
type Pricing struct {
	// Identifies the pricing model as one of "cpm", "cpc", "cpe" or "cpv".
	Model string `xml:"model,attr"`
	// The 3 letter ISO-4217 currency symbol that identifies the currency of
	// the value provided
	Currency string `xml:"currency,attr"`
	// If the value provided is to be obfuscated/encoded, publishers and advertisers
	// must negotiate the appropriate mechanism to do so. When included as part of
	// a VAST Wrapper in a chain of Wrappers, only the value offered in the first
	// Wrapper need be considered.
	Value string `xml:",cdata"`
}

type BlockedAdCategories struct {
	// Categories is a string that provides a comma separated list of category codes or labels per
	// authority that identify the ad content.
	Categories string `xml:",chardata" json:"Data"`
	// Authority is a URL for the organizational authority that produced the list being used to identify ad content.
	Authority string `xml:"authority,attr"`
}

// Wrapper element contains a URI reference to a vendor ad server (often called
// a third party ad server). The destination ad server either provides the ad
// files within a VAST <InLine> ad element or may provide a secondary Wrapper
// ad, pointing to yet another ad server. Eventually, the final ad server in
// the ad supply chain must contain all the necessary files needed to display
// the ad.
type Wrapper struct {
	// One or more URIs that directs the video player to a tracking resource file that the
	// video player should request when the first frame of the ad is displayed
	Impressions  []Impression `xml:"Impression"`
	VASTAdTagURI CDATAString
	// The name of the ad server that returned the ad
	AdSystem *AdSystem
	// Provides a value that represents a price that can be used by real-time bidding
	// (RTB) systems. VAST is not designed to handle RTB since other methods exist,
	// but this element is offered for custom solutions if needed.
	Pricing *Pricing `xml:",omitempty" json:",omitempty"`
	// A URI representing an error-tracking pixel; this element can occur multiple
	// times.
	Errors []CDATAString `xml:"Error,omitempty" json:"Error,omitempty"`
	// The ad server may provide URIs for tracking publisher-determined viewability,
	// for both the InLine ad and any Wrappers, using the <ViewableImpression> element.
	// Tracking URIs may be provided in three containers: <Viewable>, <NotViewable>, and <ViewUndetermined>.
	ViewableImpression *ViewableImpression `xml:",omitempty" json:",omitempty"`
	// The <AdVerifications> element contains one or more <Verification> elements,
	// which list the resources and metadata required to execute third-party measurement code in order to verify creative playback.
	// The <AdVerifications> element is used to contain one or more <Verification> elements,
	// which are used to initiate a controlled container where code can be executed for collecting data to verify ad playback details.
	AdVerifications []Verification `xml:",omitempty" json:",omitempty"`
	// XML node for custom extensions, as defined by the ad server. When used, a
	// custom element should be nested under <Extensions> to help separate custom
	// XML elements from VAST elements. The following example includes a custom
	// xml element within the Extensions element.
	Extensions *[]Extension `xml:"Extensions>Extension,omitempty" json:",omitempty"`
	// URL of ad tag of downstream Secondary Ad Server
	// The container for one or more <Creative> elements
	Creatives *[]CreativeWrapper `xml:"Creatives>Creative"`
	// Ad categories are used in creative separation and for compliance in certain programs. In a
	// wrapper, this field defines ad categories that cannot be returned by a downstream ad
	// server
	BlockedAdCategories *BlockedAdCategories `xml:",omitempty" json:",omitempty"`

	// Attributes
	// FollowAdditionalWrappers is a Boolean value that identifies whether subsequent Wrappers after a
	// requested VAST response is allowed. If false, any Wrappers received (i.e. not
	// an Inline VAST response) should be ignored. Otherwise, VAST Wrappers
	// received should be accepted (default value is “true.”)
	FollowAdditionalWrappers *bool `xml:"followAdditionalWrappers,attr,omitempty" json:",omitempty"`
	// AllowMultipleAds is a Boolean value that identifies whether multiple ads are allowed in the
	// requested VAST response. If true, both Pods and stand-alone ads are allowed.
	// If false, only the first stand-alone Ad (with no sequence values) in the
	// requested VAST response is allowed. Default value is “false.”
	AllowMultipleAds *bool `xml:"allowMultipleAds,attr,omitempty" json:",omitempty"`
	// FallbackOnNoAd is a Boolean value that provides instruction for using an available Ad when the
	// requested VAST response returns no ads. If true, the media player should
	// select from any stand-alone ads available. If false and the Wrapper represents
	// an Ad in a Pod, the media player should move on to the next Ad in a Pod;
	// otherwise, the media player can follow through at its own discretion where
	// no-ad responses are concerned.
	FallbackOnNoAd *bool `xml:"fallbackOnNoAd,attr,omitempty" json:",omitempty"`
}

// AdSystem contains information about the system that returned the ad
type AdSystem struct {
	// Name is a string that provides the name of the ad server that returned the ad
	Name string `xml:",chardata"`
	// Version is a string that provides the version number of the ad system that returned the ad
	Version string `xml:"version,attr,omitempty" json:"Version,omitempty"`
}

// Creative is a file that is part of a VAST ad.
type Creative struct {
	// If present, provides a VAST 4.x universal ad id
	UniversalAdID []UniversalAdID `xml:"UniversalAdId"`
	// When an API framework is needed to execute creative, a
	// CreativeExtensions> element can be added under the <Creative. This
	// extension can be used to load an executable creative with or without using
	// a media file.
	// A <CreativeExtension> element is nested under the <CreativeExtensions>
	// (plural) element so that any xml extensions are separated from VAST xml.
	// Additionally, any xml used in this extension should identify an xml name
	// space (xmlns) to avoid confusing any of the xml element names with those
	// of VAST.
	// The nested <CreativeExtension> includes an attribute for type, which
	// specifies the MIME type needed to execute the extension.
	CreativeExtensions *[]Extension `xml:"CreativeExtensions>CreativeExtension,omitempty" json:",omitempty"`
	// If defined, defines companions creatives
	CompanionAds *CompanionAds `xml:",omitempty" json:",omitempty"`
	// If present, defines a linear creative
	Linear *Linear `xml:",omitempty" json:",omitempty"`
	// If defined, defines non-linear creatives
	NonLinearAds *NonLinearAds `xml:",omitempty" json:",omitempty"`

	// Attributes

	// ID is an ad server-defined identifier for the creative
	ID string `xml:"id,attr,omitempty" json:",omitempty"`
	// Sequence is the preferred order in which multiple Creatives should be displayed
	Sequence int `xml:"sequence,attr,omitempty" json:",omitempty"`
	// AdID identifies the ad with which the creative is served
	AdID string `xml:"adId,attr,omitempty" json:",omitempty"`
	// APIFramework is the technology used for any included API
	APIFramework string `xml:"apiFramework,attr,omitempty" json:",omitempty"`
}

// CompanionAds contains companions creatives
type CompanionAds struct {
	// Provides information about which companion creative to display.
	// All means that the player must attempt to display all. Any means the player
	// must attempt to play at least one. None means all companions are optional
	Required   string      `xml:"required,attr,omitempty" json:",omitempty"`
	Companions []Companion `xml:"Companion,omitempty" json:",omitempty"`
}

// NonLinearAds contains non-linear creatives
type NonLinearAds struct {
	TrackingEvents *[]Tracking `xml:"TrackingEvents>Tracking,omitempty" json:",omitempty"`
	// Non-linear creatives
	NonLinears []NonLinear `xml:"NonLinear,omitempty" json:",omitempty"`
}

// CreativeWrapper defines wrapped creative's parent trackers
type CreativeWrapper struct {
	// An ad server-defined identifier for the creative
	ID string `xml:"id,attr,omitempty" json:",omitempty"`
	// The preferred order in which multiple Creatives should be displayed
	Sequence int `xml:"sequence,attr,omitempty" json:",omitempty"`
	// Identifies the ad with which the creative is served
	AdID string `xml:"adId,attr,omitempty" json:",omitempty"`
	// If present, defines a linear creative
	Linear *LinearWrapper `xml:",omitempty" json:",omitempty"`
	// If defined, defines companions creatives
	CompanionAds *CompanionAdsWrapper `xml:"CompanionAds,omitempty" json:",omitempty"`
	// If defined, defines non-linear creatives
	NonLinearAds *NonLinearAdsWrapper `xml:"NonLinearAds,omitempty" json:",omitempty"`
}

// CompanionAdsWrapper contains companions creatives in a wrapper
type CompanionAdsWrapper struct {
	// Provides information about which companion creative to display.
	// All means that the player must attempt to display all. Any means the player
	// must attempt to play at least one. None means all companions are optional
	Required   string             `xml:"required,attr,omitempty" json:",omitempty"`
	Companions []CompanionWrapper `xml:"Companion,omitempty" json:",omitempty"`
}

// NonLinearAdsWrapper contains non-linear creatives in a wrapper
type NonLinearAdsWrapper struct {
	TrackingEvents *[]Tracking `xml:"TrackingEvents>Tracking,omitempty" json:",omitempty"`
	// Non-linear creatives
	NonLinears []NonLinearWrapper `xml:"NonLinear,omitempty" json:",omitempty"`
}

// Linear is the most common type of video advertisement trafficked in the
// industry is a “linear ad”, which is an ad that displays in the same area
// as the content but not at the same time as the content. In fact, the video
// player must interrupt the content before displaying a linear ad.
// Linear ads are often displayed right before the video content plays.
// This ad position is called a “pre-roll” position. For this reason, a linear
// ad is often called a “pre-roll.”
type Linear struct {
	// Duration is a time value for the duration of the Linear ad in the format HH:MM:SS.mmm
	// (.mmm is optional and indicates milliseconds).
	Duration   Duration     `xml:"Duration,omitempty" json:",omitempty"`
	MediaFiles *[]MediaFile `xml:"MediaFiles>MediaFile,omitempty" json:",omitempty"`
	// AdParameters is the only way to pass information from the VAST response into the VPAID object;
	// no other mechanism is provided.
	AdParameters   *AdParameters `xml:",omitempty" json:",omitempty"`
	TrackingEvents *[]Tracking   `xml:"TrackingEvents>Tracking,omitempty" json:",omitempty"`
	VideoClicks    *VideoClicks  `xml:",omitempty" json:",omitempty"`
	Icons          *Icons        `json:",omitempty"`

	// To specify that a Linear creative can be skipped, the ad server must
	// include the skipoffset attribute in the <Linear> element. The value
	// for skipoffset is a time value in the format HH:MM:SS or HH:MM:SS.mmm
	// or a percentage in the format n%. The .mmm value in the time offset
	// represents milliseconds and is optional. This skipoffset value
	// indicates when the skip control should be provided after the creative
	// begins playing.
	SkipOffset *Offset `xml:"skipoffset,attr,omitempty" json:",omitempty"`
}

// LinearWrapper defines a wrapped linear creative
type LinearWrapper struct {
	Icons          *Icons       `json:",omitempty"`
	TrackingEvents *[]Tracking  `xml:"TrackingEvents>Tracking,omitempty" json:",omitempty"`
	VideoClicks    *VideoClicks `xml:",omitempty" json:",omitempty"`
}

// Companion defines a companion ad
type Companion struct {
	// Optional identifier
	ID string `xml:"id,attr,omitempty" json:",omitempty"`
	// Pixel dimensions of companion slot.
	Width int `xml:"width,attr,omitempty" json:",omitempty"`
	// Pixel dimensions of companion slot.
	Height int `xml:"height,attr,omitempty" json:",omitempty"`
	// Pixel dimensions of the companion asset.
	AssetWidth int `xml:"assetWidth,attr,omitempty" json:",omitempty"`
	// Pixel dimensions of the companion asset.
	AssetHeight int `xml:"assetHeight,attr,omitempty" json:",omitempty"`
	// Pixel dimensions of expanding companion ad when in expanded state.
	ExpandedWidth int `xml:"expandedWidth,attr,omitempty" json:",omitempty"`
	// Pixel dimensions of expanding companion ad when in expanded state.
	ExpandedHeight int `xml:"expandedHeight,attr,omitempty" json:",omitempty"`
	// The apiFramework defines the method to use for communication with the companion.
	APIFramework string `xml:"apiFramework,attr,omitempty" json:",omitempty"`
	// Used to match companion creative to publisher placement areas on the page.
	AdSlotID string `xml:"adSlotId,attr,omitempty" json:",omitempty"`
	// HTML to display the companion element
	HTMLResource *HTMLResource `xml:",omitempty" json:",omitempty"`
	// URL source for an IFrame to display the companion element
	IFrameResource string `xml:",omitempty" json:",omitempty"`
	// URL to a static file, such as an image or SWF file
	StaticResource *StaticResource `xml:",omitempty" json:",omitempty"`
	// Data to be passed into the companion ads. The apiFramework defines the method
	// to use for communication (e.g. “FlashVar”)
	AdParameters *AdParameters `xml:",omitempty" json:",omitempty"`
	// Alt text to be displayed when companion is rendered in HTML environment.
	AltText string `xml:",omitempty" json:",omitempty"`
	// URL to open as destination page when user clicks on the the companion banner ad.
	CompanionClickThrough string `xml:",omitempty" json:",omitempty"`
	// URLs to ping when user clicks on the companion banner ad.
	CompanionClickTrackings []CompanionClickTracking `xml:"CompanionClickTracking,omitempty" json:",omitempty"`
	// The creativeView should always be requested when present. For Companions
	// creativeView is the only supported event.
	TrackingEvents *[]Tracking `xml:"TrackingEvents>Tracking,omitempty" json:",omitempty"`
}

// CompanionWrapper defines a companion ad in a wrapper
type CompanionWrapper struct {
	// Optional identifier
	ID string `xml:"id,attr,omitempty" json:",omitempty"`
	// Pixel dimensions of companion slot.
	Width int `xml:"width,attr"`
	// Pixel dimensions of companion slot.
	Height int `xml:"height,attr"`
	// Pixel dimensions of the companion asset.
	AssetWidth int `xml:"assetWidth,attr"`
	// Pixel dimensions of the companion asset.
	AssetHeight int `xml:"assetHeight,attr"`
	// Pixel dimensions of expanding companion ad when in expanded state.
	ExpandedWidth int `xml:"expandedWidth,attr"`
	// Pixel dimensions of expanding companion ad when in expanded state.
	ExpandedHeight int `xml:"expandedHeight,attr"`
	// The apiFramework defines the method to use for communication with the companion.
	APIFramework string `xml:"apiFramework,attr,omitempty" json:",omitempty"`
	// Used to match companion creative to publisher placement areas on the page.
	AdSlotID string `xml:"adSlotId,attr,omitempty" json:",omitempty"`
	// URL to open as destination page when user clicks on the companion banner ad.
	CompanionClickThrough string `xml:",omitempty" json:",omitempty"`
	// URLs to ping when user clicks on the companion banner ad.
	CompanionClickTracking []CDATAString `xml:",omitempty" json:",omitempty"`
	// Alt text to be displayed when companion is rendered in HTML environment.
	AltText string `xml:",omitempty" json:",omitempty"`
	// The creativeView should always be requested when present. For Companions
	// creativeView is the only supported event.
	TrackingEvents *[]Tracking `xml:"TrackingEvents>Tracking,omitempty" json:",omitempty"`
	// Data to be passed into the companion ads. The apiFramework defines the method
	// to use for communication (e.g. “FlashVar”)
	AdParameters *AdParameters `xml:",omitempty" json:",omitempty"`
	// URL to a static file, such as an image or SWF file
	StaticResource *StaticResource `xml:",omitempty" json:",omitempty"`
	// URL source for an IFrame to display the companion element
	IFrameResource string `xml:",cdata"`
	// HTML to display the companion element
	HTMLResource *HTMLResource `xml:",omitempty" json:",omitempty"`
}

// NonLinear defines a non-linear ad
type NonLinear struct {
	// Optional identifier
	ID string `xml:"id,attr,omitempty" json:",omitempty"`
	// Pixel dimensions of companion.
	Width int `xml:"width,attr"`
	// Pixel dimensions of companion.
	Height int `xml:"height,attr"`
	// Pixel dimensions of expanding nonlinear ad when in expanded state.
	ExpandedWidth int `xml:"expandedWidth,attr"`
	// Pixel dimensions of expanding nonlinear ad when in expanded state.
	ExpandedHeight int `xml:"expandedHeight,attr"`
	// Whether it is acceptable to scale the image.
	Scalable *bool `xml:"scalable,attr,omitempty" json:",omitempty"`
	// Whether the ad must have its aspect ratio maintained when scales.
	MaintainAspectRatio *bool `xml:"maintainAspectRatio,attr,omitempty" json:",omitempty"`
	// Suggested duration to display non-linear ad, typically for animation to complete.
	// Expressed in standard time format hh:mm:ss.
	MinSuggestedDuration Duration `xml:"minSuggestedDuration,attr,omitempty" json:",omitempty"`
	// The apiFramework defines the method to use for communication with the nonlinear element.
	APIFramework string `xml:"apiFramework,attr,omitempty" json:",omitempty"`
	// HTML to display the companion element
	HTMLResource *HTMLResource `xml:",omitempty" json:",omitempty"`
	// URL source for an IFrame to display the companion element
	IFrameResource string `xml:",omitempty" json:",omitempty"`
	// URL to a static file, such as an image or SWF file
	StaticResource *StaticResource `xml:",omitempty" json:",omitempty"`
	// Data to be passed into the video ad.
	AdParameters *AdParameters `xml:",omitempty" json:",omitempty"`
	// URL to open as destination page when user clicks on the non-linear ad unit.
	NonLinearClickThrough *CDATAString `xml:",omitempty" json:",omitempty"`
	// URLs to ping when user clicks on the non-linear ad.
	NonLinearClickTrackings []NonLinearClickTracking `xml:"NonLinearClickTracking,omitempty" json:",omitempty"`
}

// NonLinearWrapper defines a non-linear ad in a wrapper
type NonLinearWrapper struct {
	// Optional identifier
	ID string `xml:"id,attr,omitempty" json:",omitempty"`
	// Pixel dimensions of companion.
	Width int `xml:"width,attr"`
	// Pixel dimensions of companion.
	Height int `xml:"height,attr"`
	// Pixel dimensions of expanding nonlinear ad when in expanded state.
	ExpandedWidth int `xml:"expandedWidth,attr"`
	// Pixel dimensions of expanding nonlinear ad when in expanded state.
	ExpandedHeight int `xml:"expandedHeight,attr"`
	// Whether it is acceptable to scale the image.
	Scalable *bool `xml:"scalable,attr,omitempty" json:",omitempty"`
	// Whether the ad must have its aspect ratio maintained when scales.
	MaintainAspectRatio *bool `xml:"maintainAspectRatio,attr,omitempty" json:",omitempty"`
	// Suggested duration to display non-linear ad, typically for animation to complete.
	// Expressed in standard time format hh:mm:ss.
	MinSuggestedDuration Duration `xml:"minSuggestedDuration,attr,omitempty" json:",omitempty"`
	// The apiFramework defines the method to use for communication with the nonlinear element.
	APIFramework string `xml:"apiFramework,attr,omitempty" json:",omitempty"`
	// The creativeView should always be requested when present.
	TrackingEvents *[]Tracking `xml:"TrackingEvents>Tracking,omitempty" json:",omitempty"`
	// URLs to ping when user clicks on the non-linear ad.
	NonLinearClickTracking []CDATAString `xml:",omitempty" json:",omitempty"`
}

type Icons struct {
	XMLName *xml.Name `xml:"Icons,omitempty" json:",omitempty"`
	Icon    []Icon    `xml:"Icon,omitempty" json:",omitempty"`
}

// Icon represents advertising industry initiatives like AdChoices.
type Icon struct {
	// Program identifies the industry initiative that the icon supports.
	Program string `xml:"program,attr"`
	// Width is the pixel dimensions of icon.
	Width int `xml:"width,attr"`
	// Height is the pixel dimensions of icon.
	Height int `xml:"height,attr"`
	// XPosition is the horizontal alignment location (in pixels) or a specific alignment.
	// Must match ([0-9]*|left|right)
	XPosition string `xml:"xPosition,attr"`
	// YPosition is the vertical alignment location (in pixels) or a specific alignment.
	// Must match ([0-9]*|top|bottom)
	YPosition string `xml:"yPosition,attr"`
	// Offset is the start time at which the player should display the icon. Expressed in standard time format hh:mm:ss.
	Offset *Offset `xml:"offset,attr"`
	// Duration is the time the player must display the icon. Expressed in standard time format hh:mm:ss.
	Duration Duration `xml:"duration,attr"`
	// APIFramework defines the method to use for communication with the icon element
	APIFramework string `xml:"apiFramework,attr,omitempty" json:",omitempty"`
	// Pxratio is the pixel ratio for which the icon creative is intended.
	// The pixel ratio is the ratio of physical pixels on the device to the device-independent pixels.
	// An ad intended for display on a device with a pixel ratio that is twice that of a standard 1:1 pixel ratio would use the value "2."
	// Default value is "1."
	Pxratio string `xml:"pxratio,attr,omitempty" json:",omitempty"`
	// AltText is alternative text for the image.
	// In a HTML5 image tag this should be the text for the alt attribute.
	// This should enable screen readers to properly read back a description of the icon for visually impaired users.
	AltText string `xml:"altText,attr,omitempty" json:",omitempty"`
	// HoverText is the hover text for the image.
	// In a HTML5 image tag this should be the text for the title attribute.
	HoverText string `xml:"hoverText,attr,omitempty" json:",omitempty"`

	// The view tracking for icons is used to track when the icon creative is displayed.
	// The player uses the included URI to notify the icon server when the icon has been displayed.

	// URL to a static file, such as an image or SWF file
	StaticResource *StaticResource `xml:",omitempty" json:",omitempty"`
	// URL source for an IFrame to display the companion element
	IFrameResource string `xml:",omitempty" json:",omitempty"`
	// HTML to display the companion element
	HTMLResource *HTMLResource `xml:",omitempty" json:",omitempty"`
	// URL to open as destination page when user clicks on the icon.
	IconClickThrough *CDATAString `xml:"IconClicks>IconClickThrough,omitempty" json:",omitempty"`
	// URLs to ping when user clicks on the icon.
	IconClickTrackings []CDATAString `xml:"IconClicks>IconClickTracking,omitempty" json:",omitempty"`
	// A URI for the tracking resource file to be called when the icon creative is displayed.
	IconViewTracking *CDATAString `xml:"IconViewTracking,omitempty" json:",omitempty"`
}

// Tracking defines an event tracking URL
type Tracking struct {
	// The name of the event to track for the element. The creativeView should
	// always be requested when present.
	//
	// Possible values are creativeView, start, firstQuartile, midpoint, thirdQuartile,
	// complete, mute, unmute, pause, rewind, resume, fullscreen, exitFullscreen, expand,
	// collapse, acceptInvitation, close, skip, progress.
	Event string `xml:"event,attr"`
	// The time during the video at which this url should be pinged. Must be present for
	// progress event. Must match (\d{2}:[0-5]\d:[0-5]\d(\.\d\d\d)?|1?\d?\d(\.?\d)*%)
	Offset *Offset `xml:"offset,attr,omitempty" json:",omitempty"`
	URI    string  `xml:",cdata"`

	// custom attr
	UA string `xml:"ua,attr,omitempty" json:",omitempty"`
}

// StaticResource is the URL to a static file, such as an image or SWF file
type StaticResource struct {
	// Mime type of static resource
	CreativeType string `xml:"creativeType,attr,omitempty" json:",omitempty"`
	// URL to a static file, such as an image or SWF file
	URI string `xml:",cdata"`
}

// HTMLResource is a container for HTML data
type HTMLResource struct {
	// Specifies whether the HTML is XML-encoded
	XMLEncoded *bool  `xml:"xmlEncoded,attr,omitempty" json:",omitempty"`
	HTML       string `xml:",cdata"`
}

// AdParameters defines arbitrary ad parameters
type AdParameters struct {
	// Specifies whether the parameters are XML-encoded
	XMLEncoded *bool  `xml:"xmlEncoded,attr,omitempty" json:",omitempty"`
	Parameters string `xml:",cdata"`
}

// VideoClicks contains types of video clicks
type VideoClicks struct {
	ClickTrackings []VideoClick `xml:"ClickTracking,omitempty" json:",omitempty"`
	CustomClicks   []VideoClick `xml:"CustomClick,omitempty" json:",omitempty"`
	ClickThroughs  []VideoClick `xml:"ClickThrough,omitempty" json:",omitempty"`
}

// VideoClick defines a click URL for a linear creative
type VideoClick struct {
	ID  string `xml:"id,attr,omitempty" json:",omitempty"`
	URI string `xml:",cdata"`
}

// MediaFile defines a reference to a linear creative asset
type MediaFile struct {
	// URI is a CDATA-wrapped URI to a media file.
	URI string `xml:",cdata"`

	// Attributes

	// Delivery is the method of delivery of ad (either "streaming" or "progressive")
	Delivery string `xml:"delivery,attr"`
	// Type is the MIME type. Popular MIME types include, but are not limited to
	// “video/x-ms-wmv” for Windows Media, and “video/x-flv” for Flash
	// Video. Image ads or interactive ads can be included in the
	// MediaFiles section with appropriate Mime types
	Type string `xml:"type,attr"`
	// Width is the pixel dimensions of video.
	Width int `xml:"width,attr"`
	// Height is the pixel dimensions of video.
	Height int `xml:"height,attr"`
	// Codec is the codec used to produce the media file.
	Codec string `xml:"codec,attr,omitempty" json:",omitempty"`
	// ID is an optional identifier
	ID string `xml:"id,attr,omitempty" json:",omitempty"`
	// Bitrate of encoded video in Kbps. If bitrate is supplied, MinBitrate
	// and MaxBitrate should not be supplied.
	Bitrate int `xml:"bitrate,attr,omitempty" json:",omitempty"`
	// MinBitrate is the minimum bitrate of an adaptive stream in Kbps. If MinBitrate is supplied,
	// MaxBitrate must be supplied and Bitrate should not be supplied.
	MinBitrate int `xml:"minBitrate,attr,omitempty" json:",omitempty"`
	// MaxBitrate is the maximum bitrate of an adaptive stream in Kbps. If MaxBitrate is supplied,
	// MinBitrate must be supplied and Bitrate should not be supplied.
	MaxBitrate int `xml:"maxBitrate,attr,omitempty" json:",omitempty"`
	// Scalable determines whether it is acceptable to scale the image.
	Scalable *bool `xml:"scalable,attr,omitempty" json:",omitempty"`
	// MaintainAspectRatio determines whether the ad must have its aspect ratio maintained when scales.
	MaintainAspectRatio *bool `xml:"maintainAspectRatio,attr,omitempty" json:",omitempty"`
	// APIFramework is the APIFramework defines the method to use for communication if the MediaFile
	// is interactive. Suggested values for this element are “VPAID”, “FlashVars”
	// (for Flash/Flex), “initParams” (for Silverlight) and “GetVariables” (variables
	// placed in key/value pairs on the asset request).
	APIFramework string `xml:"apiFramework,attr,omitempty" json:",omitempty"`
	// FileSize is an optional field that helps eliminate the need to calculate the size based on bitrate and duration.
	FileSize int `xml:"fileSize,attr,omitempty" json:",omitempty"`
	// MediaType is the type of media file (2D / 3D / 360 / etc).
	MediaType string `xml:"mediaType,attr,omitempty" json:",omitempty"`
}

// UniversalAdID describes a VAST 4.x universal ad id.
type UniversalAdID struct {
	// ID is a string identifying the unique creative identifier. Default value is “unknown”.
	ID string `xml:",chardata" json:"Data"`
	// IDRegistry is a string used to identify the URL for the registry website where the unique
	// creative ID is cataloged. Default value is “unknown.”
	IDRegistry string `xml:"idRegistry,attr"`
}

// Category
// Used in creative separation and for compliance in certain programs, a category field is
// needed to categorize the ad’s content. Several category lists exist, some for describing site
// content and some for describing ad content. Some lists are used interchangeably for both
// site content and ad content. For example, the category list used to comply with the IAB
// Quality Assurance Guidelines (QAG) describes site content, but is sometimes used to
// describe ad content.
// The VAST category field should only use AD CONTENT description categories.
// The authority attribute is used to identify the organizational authority that developed the
// list being used. In some cases, the publisher may require that an ad category be identified.
// If required by the publisher and not provided, the publisher may skip the ad, notify the ad
// server using the <Error> URI, if provided (error code 204), and move on to the next option.
// If category is used, the authority= attribute must be provided.
type Category struct {
	// Category is a string that provides the name of the ad server that returned the ad
	Category string `xml:",chardata" json:"Data"`
	// Authority is a URL for the organizational authority that produced the list being used to identify
	// ad content category
	Authority string `xml:"authority,attr"`
}

// Survey
// Ad tech vendors may want to use the ad to collect data for resource purposes. The
// <Survey> element can be used to provide a URI to any resource file having to do with
// collecting survey data. Publishers and any parties using the <Survey> element should
// determine how surveys are implemented and executed. Multiple survey elements may be
// provided.
// A type attribute is available to specify the MIME type being served. For example, the
// attribute might be set to type="text/javascript". Surveys can be dynamically inserted
// into the VAST response as long as cross-domain issues are avoided
// Deprecated - VAST 4.1: Since usage was very limited and survey implementations can be
// supported by other VAST elements such as 3rd party trackers.
type Survey struct {
	// URI is a URI to any resource relating to an integrated survey.
	URI string `xml:",cdata"`
	// Type is the MIME type of the resource being served.
	Type string `xml:"type,attr"`
}

type MediaFiles struct {
	MediaFile               []MediaFile
	Mezzanine               []Mezzanine               `xml:",omitempty" json:",omitempty"`
	InteractiveCreativeFile []InteractiveCreativeFile `xml:",omitempty" json:",omitempty"`
	ClosedCaptionFiles      *[]ClosedCaptionFile      `xml:"ClosedCaptionFiles>ClosedCaptionFile,omitempty" json:",omitempty"`
}

// CompanionClickTracking element is used to track the click
type CompanionClickTracking struct {
	// An id provided by the ad server to track the click in reports.
	ID  string `xml:"id,attr,omitempty" json:",omitempty"`
	URI string `xml:",cdata"`
}

// NonLinearClickTracking element is used to track the click
type NonLinearClickTracking struct {
	// An id provided by the ad server to track the click in reports
	ID  string `xml:"id,attr,omitempty" json:",omitempty"`
	URI string `xml:",cdata"`
}
