package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Status int  `json:"status"`
	Data   Data `json:"data"`
}
type URL struct {
	Base any    `json:"base"`
	URI  string `json:"uri"`
}
type DataLayerP struct {
	Brand         string  `json:"brand"`
	Category      string  `json:"category"`
	Metric6       int     `json:"metric6"`
	Dimension2    float32 `json:"dimension2"`
	Dimension6    float32 `json:"dimension6"`
	Dimension7    string  `json:"dimension7"`
	Dimension9    float32 `json:"dimension9"`
	Dimension11   float32 `json:"dimension11"`
	Dimension20   string  `json:"dimension20"`
	ItemCategory2 string  `json:"item_category2"`
	ItemCategory3 string  `json:"item_category3"`
	ItemCategory4 string  `json:"item_category4"`
	ItemCategory5 string  `json:"item_category5"`
}
type ServiceList struct {
	Title string `json:"title"`
}
type Digiplus struct {
	Services                     []string      `json:"services"`
	ServicesSummary              []string      `json:"services_summary"`
	ServiceList                  []ServiceList `json:"service_list"`
	IsJetEligible                bool          `json:"is_jet_eligible"`
	CashBack                     int           `json:"cash_back"`
	IsGeneralLocationJetEligible bool          `json:"is_general_location_jet_eligible"`
	FastShippingText             any           `json:"fast_shipping_text"`
}
type Main struct {
	StorageIds   []any    `json:"storage_ids"`
	URL          []string `json:"url"`
	ThumbnailURL any      `json:"thumbnail_url"`
	TemporaryID  any      `json:"temporary_id"`
	WebpURL      []string `json:"webp_url"`
}
type List struct {
	StorageIds   []any    `json:"storage_ids"`
	URL          []string `json:"url"`
	ThumbnailURL any      `json:"thumbnail_url"`
	TemporaryID  any      `json:"temporary_id"`
	WebpURL      []string `json:"webp_url"`
}
type Images struct {
	Main Main   `json:"main"`
	List []List `json:"list"`
}
type RatingProd struct {
	Rate  float32 `json:"rate"`
	Count int     `json:"count"`
}

type Warranty struct {
	ID      int    `json:"id"`
	TitleFa string `json:"title_fa"`
	TitleEn string `json:"title_en"`
}
type Size struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
type Rating struct {
	TotalRate      int     `json:"total_rate"`
	TotalCount     int     `json:"total_count"`
	Commitment     float64 `json:"commitment"`
	NoReturn       float64 `json:"no_return"`
	OnTimeShipping float64 `json:"on_time_shipping"`
}
type PropertiesSeller struct {
	IsTrusted  bool `json:"is_trusted"`
	IsOfficial bool `json:"is_official"`
	IsRoosta   bool `json:"is_roosta"`
	IsNew      bool `json:"is_new"`
}
type Grade struct {
	Label string `json:"label"`
	Color string `json:"color"`
}
type Seller struct {
	ID               int              `json:"id"`
	Title            string           `json:"title"`
	Code             string           `json:"code"`
	URL              string           `json:"url"`
	Rating           Rating           `json:"rating"`
	Properties       PropertiesSeller `json:"properties"`
	Stars            float64          `json:"stars"`
	Grade            Grade            `json:"grade"`
	Logo             any              `json:"logo"`
	RegistrationDate string           `json:"registration_date"`
}
type Digiclub struct {
	Point int `json:"point"`
}
type GoldPriceDetails struct {
	GoldPrice        int     `json:"gold_price"`
	NonGoldPartsCost int     `json:"non_gold_parts_cost"`
	NonGoldPartsWage int     `json:"non_gold_parts_wage"`
	GoldWage         int     `json:"gold_wage"`
	Profit           int     `json:"profit"`
	Tax              int     `json:"tax"`
	Size             float64 `json:"size"`
	Carat            string  `json:"carat"`
	GoldTotalPrice   int     `json:"gold_total_price"`
	TotalProfit      float64 `json:"total_profit"`
	TotalTax         float64 `json:"total_tax"`
	TotalGoldWage    int     `json:"total_gold_wage"`
}
type Price struct {
	SellingPrice        int              `json:"selling_price"`
	RrpPrice            int              `json:"rrp_price"`
	OrderLimit          int              `json:"order_limit"`
	IsIncredible        bool             `json:"is_incredible"`
	IsPromotion         bool             `json:"is_promotion"`
	IsLockedForDigiplus bool             `json:"is_locked_for_digiplus"`
	BnplActive          bool             `json:"bnpl_active"`
	DiscountPercent     int              `json:"discount_percent"`
	GoldPriceDetails    GoldPriceDetails `json:"gold_price_details"`
	IsPlusEarlyAccess   bool             `json:"is_plus_early_access"`
}
type Providers struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	HasLeadTime bool   `json:"has_lead_time"`
	Type        string `json:"type"`
}
type ShipmentMethods struct {
	Description string      `json:"description"`
	HasLeadTime bool        `json:"has_lead_time"`
	Providers   []Providers `json:"providers"`
}
type DefaultVariant struct {
	ID                       int             `json:"id"`
	LeadTime                 int             `json:"lead_time"`
	Rank                     float64         `json:"rank"`
	Rate                     int             `json:"rate"`
	Statistics               any             `json:"statistics"`
	Status                   string          `json:"status"`
	Properties               Properties      `json:"properties"`
	Digiplus                 Digiplus        `json:"digiplus"`
	Warranty                 Warranty        `json:"warranty"`
	Size                     Size            `json:"size"`
	Seller                   Seller          `json:"seller"`
	Digiclub                 Digiclub        `json:"digiclub"`
	Price                    Price           `json:"price"`
	ShipmentMethods          ShipmentMethods `json:"shipment_methods"`
	HasImporterPrice         bool            `json:"has_importer_price"`
	ManufacturePriceNotExist bool            `json:"manufacture_price_not_exist"`
	HasBestPriceInLastMonth  bool            `json:"has_best_price_in_last_month"`
	BuyBoxNotices            []any           `json:"buy_box_notices"`
	VariantBadges            []any           `json:"variant_badges"`
}
type PropertiesProduct struct {
	IsFastShipping      bool  `json:"is_fast_shipping"`
	IsShipBySeller      bool  `json:"is_ship_by_seller"`
	FreeShippingBadge   bool  `json:"free_shipping_badge"`
	IsMultiWarehouse    bool  `json:"is_multi_warehouse"`
	IsFake              bool  `json:"is_fake"`
	HasGift             bool  `json:"has_gift"`
	MinPriceInLastMonth int   `json:"min_price_in_last_month"`
	IsNonInventory      bool  `json:"is_non_inventory"`
	IsAd                bool  `json:"is_ad"`
	Ad                  []any `json:"ad"`
	IsJetEligible       bool  `json:"is_jet_eligible"`
	IsMedicalSupplement bool  `json:"is_medical_supplement"`
	HasPrintedPrice     bool  `json:"has_printed_price"`
}
type Category struct {
	ID                 int    `json:"id"`
	TitleFa            string `json:"title_fa"`
	TitleEn            string `json:"title_en"`
	Code               string `json:"code"`
	ContentDescription string `json:"content_description"`
	ContentBox         string `json:"content_box"`
	ReturnReasonAlert  string `json:"return_reason_alert"`
}
type Logo struct {
	StorageIds   []any    `json:"storage_ids"`
	URL          []string `json:"url"`
	ThumbnailURL any      `json:"thumbnail_url"`
	TemporaryID  any      `json:"temporary_id"`
	WebpURL      any      `json:"webp_url"`
}
type Brand struct {
	ID              int    `json:"id"`
	Code            string `json:"code"`
	TitleFa         string `json:"title_fa"`
	TitleEn         string `json:"title_en"`
	URL             URL    `json:"url"`
	Visibility      bool   `json:"visibility"`
	Logo            Logo   `json:"logo"`
	IsPremium       bool   `json:"is_premium"`
	IsMiscellaneous bool   `json:"is_miscellaneous"`
	IsNameSimilar   bool   `json:"is_name_similar"`
}
type Attributes struct {
	Title  string   `json:"title"`
	Values []string `json:"values"`
}
type Review struct {
	Description string       `json:"description"`
	Attributes  []Attributes `json:"attributes"`
}
type ProsAndCons struct {
	Advantages    []any `json:"advantages"`
	Disadvantages []any `json:"disadvantages"`
}
type Suggestion struct {
	Count      int `json:"count"`
	Percentage int `json:"percentage"`
}
type Properties struct {
	IsFastShipping      bool `json:"is_fast_shipping"`
	IsShipBySeller      bool `json:"is_ship_by_seller"`
	IsMultiWarehouse    bool `json:"is_multi_warehouse"`
	HasSimilarVariants  bool `json:"has_similar_variants"`
	IsRural             bool `json:"is_rural"`
	InDigikalaWarehouse bool `json:"in_digikala_warehouse"`
}

type Variants struct {
	ID                       int             `json:"id"`
	LeadTime                 int             `json:"lead_time"`
	Rank                     float64         `json:"rank"`
	Rate                     int             `json:"rate"`
	Statistics               any             `json:"statistics"`
	Status                   string          `json:"status"`
	Properties               Properties      `json:"properties"`
	Digiplus                 Digiplus        `json:"digiplus"`
	Warranty                 Warranty        `json:"warranty"`
	Size                     Size            `json:"size"`
	Seller                   Seller          `json:"seller"`
	Digiclub                 Digiclub        `json:"digiclub"`
	Price                    Price           `json:"price"`
	ShipmentMethods          ShipmentMethods `json:"shipment_methods"`
	HasImporterPrice         bool            `json:"has_importer_price"`
	ManufacturePriceNotExist bool            `json:"manufacture_price_not_exist"`
	HasBestPriceInLastMonth  bool            `json:"has_best_price_in_last_month"`
	BuyBoxNotices            []any           `json:"buy_box_notices"`
	VariantBadges            []any           `json:"variant_badges"`
}
type SecondDefaultVariant struct {
	ID                       int             `json:"id"`
	LeadTime                 int             `json:"lead_time"`
	Rank                     float64         `json:"rank"`
	Rate                     int             `json:"rate"`
	Statistics               any             `json:"statistics"`
	Status                   string          `json:"status"`
	Properties               Properties      `json:"properties"`
	Digiplus                 Digiplus        `json:"digiplus"`
	Warranty                 Warranty        `json:"warranty"`
	Size                     Size            `json:"size"`
	Seller                   Seller          `json:"seller"`
	Digiclub                 Digiclub        `json:"digiclub"`
	Price                    Price           `json:"price"`
	ShipmentMethods          ShipmentMethods `json:"shipment_methods"`
	HasImporterPrice         bool            `json:"has_importer_price"`
	ManufacturePriceNotExist bool            `json:"manufacture_price_not_exist"`
	HasBestPriceInLastMonth  bool            `json:"has_best_price_in_last_month"`
	BuyBoxNotices            []any           `json:"buy_box_notices"`
	VariantBadges            []any           `json:"variant_badges"`
}
type Breadcrumb struct {
	Title string `json:"title"`
	URL   URL    `json:"url"`
}
type Specifications struct {
	Title      string       `json:"title"`
	Attributes []Attributes `json:"attributes"`
}
type TechnicalProperties struct {
	Advantages    []any `json:"advantages"`
	Disadvantages []any `json:"disadvantages"`
}
type ExpertReviews struct {
	Attributes          []any               `json:"attributes"`
	Description         string              `json:"description"`
	ShortReview         string              `json:"short_review"`
	AdminRates          []any               `json:"admin_rates"`
	ReviewSections      []any               `json:"review_sections"`
	TechnicalProperties TechnicalProperties `json:"technical_properties"`
}
type BrandCategoryURL struct {
	Base any    `json:"base"`
	URI  string `json:"uri"`
}
type Meta struct {
	BrandCategoryURL BrandCategoryURL `json:"brand_category_url"`
}
type StCmpTacker struct {
	Neo    string `json:"neo"`
	Cx     string `json:"cx"`
	Dx     string `json:"dx"`
	DataFx string `json:"data-fx"`
	Zero   string `json:"zero"`
}
type Product struct {
	ID                   int                  `json:"id"`
	TitleFa              string               `json:"title_fa"`
	TitleEn              string               `json:"title_en"`
	URL                  URL                  `json:"url"`
	Status               string               `json:"status"`
	HasQuickView         bool                 `json:"has_quick_view"`
	DataLayer            DataLayerP           `json:"data_layer"`
	ProductType          string               `json:"product_type"`
	TestTitleFa          string               `json:"test_title_fa"`
	TestTitleEn          string               `json:"test_title_en"`
	Digiplus             Digiplus             `json:"digiplus"`
	Images               Images               `json:"images"`
	Rating               RatingProd           `json:"rating"`
	Colors               []any                `json:"colors"`
	DefaultVariant       DefaultVariant       `json:"default_variant"`
	Properties           PropertiesProduct    `json:"properties"`
	Badges               []any                `json:"badges"`
	HasTrueToSize        bool                 `json:"has_true_to_size"`
	ProductBadges        []any                `json:"product_badges"`
	Videos               []any                `json:"videos"`
	Category             Category             `json:"category"`
	Brand                Brand                `json:"brand"`
	Review               Review               `json:"review"`
	ProsAndCons          ProsAndCons          `json:"pros_and_cons"`
	Suggestion           Suggestion           `json:"suggestion"`
	Variants             []Variants           `json:"variants"`
	SecondDefaultVariant SecondDefaultVariant `json:"second_default_variant"`
	QuestionsCount       int                  `json:"questions_count"`
	CommentsCount        int                  `json:"comments_count"`
	CommentsOverview     json.RawMessage    `json:"comments_overview,omitempty"`
	Breadcrumb           []Breadcrumb         `json:"breadcrumb"`
	HasSizeGuide         bool                 `json:"has_size_guide"`
	Specifications       []Specifications     `json:"specifications"`
	ExpertReviews        ExpertReviews        `json:"expert_reviews"`
	Meta                 Meta                 `json:"meta"`
	LastComments         []any                `json:"last_comments"`
	LastQuestions        []any                `json:"last_questions"`
	Tags                 []any                `json:"tags"`
	DigifyTouchpoint     string               `json:"digify_touchpoint"`
	ShowType             string               `json:"show_type"`
	StCmpTacker          StCmpTacker          `json:"st_cmp_tacker"`

	IsInactive bool `json:"is_inactive,omitempty"`
}
 
type CommentsOverview struct {
	ID            int      `json:"id"`
	Overview      string   `json:"overview"`
	Advantages    []string `json:"advantages"`
	Disadvantages []string `json:"disadvantages"`
}

type ActionField struct {
	List string `json:"list"`
}
type Products struct {
	Brand         string  `json:"brand"`
	Category      string  `json:"category"`
	Metric6       int     `json:"metric6"`
	Dimension2    float32 `json:"dimension2"`
	Dimension6    float32 `json:"dimension6"`
	Dimension7    string  `json:"dimension7"`
	Dimension9    float32 `json:"dimension9"`
	Dimension11   float32 `json:"dimension11"`
	Dimension20   string  `json:"dimension20"`
	ItemCategory2 string  `json:"item_category2"`
	ItemCategory3 string  `json:"item_category3"`
	ItemCategory4 string  `json:"item_category4"`
	ItemCategory5 string  `json:"item_category5"`
	Name          string  `json:"name"`
	ID            int     `json:"id"`
	Price         int     `json:"price"`
	Metric8       int     `json:"metric8"`
	Dimension3    string  `json:"dimension3"`
	Dimension10   int     `json:"dimension10"`
	Dimension15   int     `json:"dimension15"`
	Metric15      int     `json:"metric15"`
	Metric11      int     `json:"metric11"`
	Metric12      int     `json:"metric12"`
}
type Detail struct {
	ActionField ActionField `json:"actionField"`
	Products    []Products  `json:"products"`
}
type Ecommerce struct {
	Detail Detail `json:"detail"`
}
type DataLayer struct {
	Event     string    `json:"event"`
	Ecommerce Ecommerce `json:"ecommerce"`
}
type TwitterCard struct {
	Title       string `json:"title"`
	Image       string `json:"image"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}
type OpenGraph struct {
	Title        string `json:"title"`
	URL          string `json:"url"`
	Image        string `json:"image"`
	Availability string `json:"availability"`
	Type         string `json:"type"`
	Site         string `json:"site"`
	Price        int    `json:"price"`
	Description  any    `json:"description"`
}
type Header struct {
	Title        string `json:"title"`
	Description  string `json:"description"`
	CanonicalURL string `json:"canonical_url"`
}
type Seo struct {
	Title        string      `json:"title"`
	Description  string      `json:"description"`
	TwitterCard  TwitterCard `json:"twitter_card"`
	OpenGraph    OpenGraph   `json:"open_graph"`
	Header       Header      `json:"header"`
	MarkupSchema []any       `json:"markup_schema"`
}
type EventData struct {
	Currency        string   `json:"currency"`
	DeviceType      string   `json:"deviceType"`
	Name            string   `json:"name"`
	ProductID       int      `json:"productId"`
	ProductImageURL []string `json:"productImageUrl"`
	Quantity        any      `json:"quantity"`
	LeafCategory    string   `json:"leafCategory"`
	UnitPrice       int      `json:"unitPrice"`
	URL             string   `json:"url"`
	SupplyCategory  string   `json:"supplyCategory"`
	CategoryLevel1  string   `json:"categoryLevel1"`
	CategoryLevel2  string   `json:"categoryLevel2"`
	CategoryLevel3  string   `json:"categoryLevel3"`
	CategoryLevel4  string   `json:"categoryLevel4"`
	CategoryLevel5  string   `json:"categoryLevel5"`
}
type Intrack struct {
	EventName string    `json:"eventName"`
	EventData EventData `json:"eventData"`
	UserID    any       `json:"userId"`
}
type PromotionBanner struct {
	DkjetPromotion any `json:"dkjet_promotion"`
}
type PageInfo struct {
	ProductID int `json:"product_id"`
}
type BigdataTrackerData struct {
	PageName string   `json:"page_name"`
	PageInfo PageInfo `json:"page_info"`
}
type Data struct {
	Product            Product            `json:"product"`
	DataLayer          DataLayer          `json:"data_layer"`
	Seo                Seo                `json:"seo"`
	Intrack            Intrack            `json:"intrack"`
	LandingTouchpoint  []any              `json:"landing_touchpoint"`
	DynamicTouchPoints []any              `json:"dynamic_touch_points"`
	PromotionBanner    PromotionBanner    `json:"promotion_banner"`
	BigdataTrackerData BigdataTrackerData `json:"bigdata_tracker_data"`
}

func main() {
	url := "https://api.digikala.com/v2/product/16175190/"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	err = ioutil.WriteFile("response.json", body, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
	 

	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	if response.Data.Product.IsInactive {
		fmt.Println("Product is inactive.")
	} else {

		fmt.Println("Product ID:", response.Data.Product.ID)
		fmt.Println("Title (FA):", response.Data.Product.TitleFa)
		fmt.Println("Category:", response.Data.Product.DataLayer.Category)

	}
}
