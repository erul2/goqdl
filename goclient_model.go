package goqdl

import (
	"fmt"
	"strings"
)

type GroupIds []int

type CredParameters struct {
	ShortLabel string `json:"short_label"`
}

func (p CredParameters) IsEmpty() bool {
	return p == CredParameters{}
}

type LoginResponse struct {
	User struct {
		Id           int         `json:"id"`
		PublicId     string      `json:"publicId"`
		Email        string      `json:"email"`
		Login        string      `json:"login"`
		Firstname    interface{} `json:"firstname"`
		Lastname     interface{} `json:"lastname"`
		DisplayName  string      `json:"display_name"`
		CountryCode  string      `json:"country_code"`
		LanguageCode string      `json:"language_code"`
		Zone         string      `json:"zone"`
		Store        string      `json:"store"`
		Country      string      `json:"country"`
		Avatar       string      `json:"avatar"`
		Genre        string      `json:"genre"`
		Age          int         `json:"age"`
		CreationDate string      `json:"creation_date"`
		Subscription struct {
			Offer            string `json:"offer"`
			Periodicity      string `json:"periodicity"`
			StartDate        string `json:"start_date"`
			EndDate          string `json:"end_date"`
			IsCanceled       bool   `json:"is_canceled"`
			HouseholdSizeMax int    `json:"household_size_max"`
		} `json:"subscription"`
		Credential struct {
			Id          int    `json:"id"`
			Label       string `json:"label"`
			Description string `json:"description"`
			Parameters  struct {
				LossyStreaming          bool  `json:"lossy_streaming"`
				LosslessStreaming       bool  `json:"lossless_streaming"`
				HiresStreaming          bool  `json:"hires_streaming"`
				HiresPurchasesStreaming bool  `json:"hires_purchases_streaming"`
				MobileStreaming         bool  `json:"mobile_streaming"`
				OfflineStreaming        bool  `json:"offline_streaming"`
				HfpPurchase             bool  `json:"hfp_purchase"`
				IncludedFormatGroupIds  []int `json:"included_format_group_ids"`
				ColorScheme             struct {
					Logo string `json:"logo"`
				} `json:"color_scheme"`
				Label      string `json:"label"`
				ShortLabel string `json:"short_label"`
				Source     string `json:"source"`
			} `json:"parameters"`
		} `json:"credential"`
		LastUpdate struct {
			Favorite       int `json:"favorite"`
			FavoriteAlbum  int `json:"favorite_album"`
			FavoriteArtist int `json:"favorite_artist"`
			FavoriteTrack  int `json:"favorite_track"`
			Playlist       int `json:"playlist"`
			Purchase       int `json:"purchase"`
		} `json:"last_update"`
		StoreFeatures struct {
			Download                 bool `json:"download"`
			Streaming                bool `json:"streaming"`
			Editorial                bool `json:"editorial"`
			Club                     bool `json:"club"`
			Wallet                   bool `json:"wallet"`
			Weeklyq                  bool `json:"weeklyq"`
			Autoplay                 bool `json:"autoplay"`
			InappPurchaseSubscripton bool `json:"inapp_purchase_subscripton"`
			OptIn                    bool `json:"opt_in"`
			MusicImport              bool `json:"music_import"`
		} `json:"store_features"`
		PlayerSettings struct {
			SonosAudioFormat int `json:"sonos_audio_format"`
		} `json:"player_settings"`
		Externals struct {
		} `json:"externals"`
	} `json:"user"`
	UserAuthToken string `json:"user_auth_token"`
}

type GetFileURLResponse struct {
	TrackId      int    `json:"track_id"`
	Duration     int    `json:"duration"`
	Url          string `json:"url"`
	FormatId     int    `json:"format_id"`
	MimeType     string `json:"mime_type"`
	Sample       bool   `json:"sample"`
	Restrictions []struct {
		Code string `json:"code"`
	} `json:"restrictions"`
	SamplingRate float64 `json:"sampling_rate"`
	BitDepth     int     `json:"bit_depth"`
}

type GetTrackMetaResponse struct {
	MaximumBitDepth int    `json:"maximum_bit_depth"`
	Copyright       string `json:"copyright"`
	Performers      string `json:"performers"`
	AudioInfo       struct {
		ReplaygainTrackGain float64 `json:"replaygain_track_gain"`
		ReplaygainTrackPeak float64 `json:"replaygain_track_peak"`
	} `json:"audio_info"`
	Performer struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"performer"`
	Album struct {
		MaximumBitDepth int `json:"maximum_bit_depth"`
		Image           struct {
			Small     string      `json:"small"`
			Thumbnail string      `json:"thumbnail"`
			Large     string      `json:"large"`
			Back      interface{} `json:"back"`
		} `json:"image"`
		MediaCount int `json:"media_count"`
		Artist     struct {
			Image       interface{} `json:"image"`
			Name        string      `json:"name"`
			Id          int         `json:"id"`
			AlbumsCount int         `json:"albums_count"`
			Slug        string      `json:"slug"`
			Picture     interface{} `json:"picture"`
		} `json:"artist"`
		Artists []struct {
			Id    int      `json:"id"`
			Name  string   `json:"name"`
			Roles []string `json:"roles"`
		} `json:"artists"`
		Upc        string `json:"upc"`
		ReleasedAt int    `json:"released_at"`
		Label      struct {
			Name        string `json:"name"`
			Id          int    `json:"id"`
			AlbumsCount int    `json:"albums_count"`
			SupplierId  int    `json:"supplier_id"`
			Slug        string `json:"slug"`
		} `json:"label"`
		Title           string      `json:"title"`
		QobuzId         int         `json:"qobuz_id"`
		Version         interface{} `json:"version"`
		Url             string      `json:"url"`
		Duration        int         `json:"duration"`
		ParentalWarning bool        `json:"parental_warning"`
		Popularity      int         `json:"popularity"`
		TracksCount     int         `json:"tracks_count"`
		Genre           struct {
			Path  []int  `json:"path"`
			Color string `json:"color"`
			Name  string `json:"name"`
			Id    int    `json:"id"`
			Slug  string `json:"slug"`
		} `json:"genre"`
		MaximumChannelCount            int           `json:"maximum_channel_count"`
		Id                             string        `json:"id"`
		MaximumSamplingRate            float64       `json:"maximum_sampling_rate"`
		Articles                       []interface{} `json:"articles"`
		ReleaseDateOriginal            string        `json:"release_date_original"`
		ReleaseDateDownload            string        `json:"release_date_download"`
		ReleaseDateStream              string        `json:"release_date_stream"`
		Purchasable                    bool          `json:"purchasable"`
		Streamable                     bool          `json:"streamable"`
		Previewable                    bool          `json:"previewable"`
		Sampleable                     bool          `json:"sampleable"`
		Downloadable                   bool          `json:"downloadable"`
		Displayable                    bool          `json:"displayable"`
		PurchasableAt                  int           `json:"purchasable_at"`
		StreamableAt                   int           `json:"streamable_at"`
		Hires                          bool          `json:"hires"`
		HiresStreamable                bool          `json:"hires_streamable"`
		Awards                         []interface{} `json:"awards"`
		Goodies                        []interface{} `json:"goodies"`
		Area                           interface{}   `json:"area"`
		Catchline                      string        `json:"catchline"`
		CreatedAt                      int           `json:"created_at"`
		GenresList                     []string      `json:"genres_list"`
		Period                         interface{}   `json:"period"`
		Copyright                      string        `json:"copyright"`
		IsOfficial                     bool          `json:"is_official"`
		MaximumTechnicalSpecifications string        `json:"maximum_technical_specifications"`
		ProductSalesFactorsMonthly     int           `json:"product_sales_factors_monthly"`
		ProductSalesFactorsWeekly      int           `json:"product_sales_factors_weekly"`
		ProductSalesFactorsYearly      int           `json:"product_sales_factors_yearly"`
		ProductType                    string        `json:"product_type"`
		ProductUrl                     string        `json:"product_url"`
		RecordingInformation           string        `json:"recording_information"`
		RelativeUrl                    string        `json:"relative_url"`
		ReleaseTags                    []interface{} `json:"release_tags"`
		ReleaseType                    string        `json:"release_type"`
		Slug                           string        `json:"slug"`
		Subtitle                       string        `json:"subtitle"`
		Description                    string        `json:"description"`
	} `json:"album"`
	Work                interface{}   `json:"work"`
	Isrc                string        `json:"isrc"`
	Title               string        `json:"title"`
	Version             string        `json:"version"`
	Duration            int           `json:"duration"`
	ParentalWarning     bool          `json:"parental_warning"`
	TrackNumber         int           `json:"track_number"`
	MaximumChannelCount int           `json:"maximum_channel_count"`
	Id                  int           `json:"id"`
	MediaNumber         int           `json:"media_number"`
	MaximumSamplingRate float64       `json:"maximum_sampling_rate"`
	Articles            []interface{} `json:"articles"`
	ReleaseDateOriginal interface{}   `json:"release_date_original"`
	ReleaseDateDownload interface{}   `json:"release_date_download"`
	ReleaseDateStream   interface{}   `json:"release_date_stream"`
	ReleaseDatePurchase interface{}   `json:"release_date_purchase"`
	Purchasable         bool          `json:"purchasable"`
	Streamable          bool          `json:"streamable"`
	Previewable         bool          `json:"previewable"`
	Sampleable          bool          `json:"sampleable"`
	Downloadable        bool          `json:"downloadable"`
	Displayable         bool          `json:"displayable"`
	PurchasableAt       interface{}   `json:"purchasable_at"`
	StreamableAt        int           `json:"streamable_at"`
	Hires               bool          `json:"hires"`
	HiresStreamable     bool          `json:"hires_streamable"`
}

func (m GetTrackMetaResponse) GetTitle() string {
	title := m.Title
	if m.Version != "" {
		title += " (" + m.Version + ")"
	}
	return title
}

type TrackAttributes struct {
	Album        string
	Artist       string
	TrackTitle   string
	Year         string
	BitDepth     int
	SamplingRate float64
}

func (m GetTrackMetaResponse) GetTrackAttributes(fileUrl GetFileURLResponse) TrackAttributes {
	return TrackAttributes{
		Album:        m.Album.Title,
		Artist:       m.Album.Artist.Name,
		TrackTitle:   m.GetTitle(),
		Year:         m.Album.ReleaseDateOriginal,
		BitDepth:     fileUrl.BitDepth,
		SamplingRate: fileUrl.SamplingRate,
	}
}

func (attributes TrackAttributes) folderFormat() string {
	year := strings.Split(attributes.Year, "-")[0]
	return fmt.Sprintf("%s - %s (%s) [%dB-%.1fkHz]", attributes.Artist, attributes.Album, year, attributes.BitDepth, attributes.SamplingRate)
}

type GetAlbumMetaResponse struct {
	MaximumBitDepth int `json:"maximum_bit_depth"`
	Image           struct {
		Small     string      `json:"small"`
		Thumbnail string      `json:"thumbnail"`
		Large     string      `json:"large"`
		Back      interface{} `json:"back"`
	} `json:"image"`
	MediaCount int `json:"media_count"`
	Artist     struct {
		Image       interface{} `json:"image"`
		Name        string      `json:"name"`
		Id          int         `json:"id"`
		AlbumsCount int         `json:"albums_count"`
		Slug        string      `json:"slug"`
		Picture     interface{} `json:"picture"`
	} `json:"artist"`
	Artists []struct {
		Id    int      `json:"id"`
		Name  string   `json:"name"`
		Roles []string `json:"roles"`
	} `json:"artists"`
	Upc        string `json:"upc"`
	ReleasedAt int    `json:"released_at"`
	Label      struct {
		Name        string `json:"name"`
		Id          int    `json:"id"`
		AlbumsCount int    `json:"albums_count"`
		SupplierId  int    `json:"supplier_id"`
		Slug        string `json:"slug"`
	} `json:"label"`
	Title           string `json:"title"`
	QobuzId         int    `json:"qobuz_id"`
	Version         string `json:"version"`
	Url             string `json:"url"`
	Duration        int    `json:"duration"`
	ParentalWarning bool   `json:"parental_warning"`
	Popularity      int    `json:"popularity"`
	TracksCount     int    `json:"tracks_count"`
	Genre           struct {
		Path  []int  `json:"path"`
		Color string `json:"color"`
		Name  string `json:"name"`
		Id    int    `json:"id"`
		Slug  string `json:"slug"`
	} `json:"genre"`
	MaximumChannelCount int           `json:"maximum_channel_count"`
	Id                  string        `json:"id"`
	MaximumSamplingRate float64       `json:"maximum_sampling_rate"`
	Articles            []interface{} `json:"articles"`
	ReleaseDateOriginal string        `json:"release_date_original"`
	ReleaseDateDownload string        `json:"release_date_download"`
	ReleaseDateStream   string        `json:"release_date_stream"`
	Purchasable         bool          `json:"purchasable"`
	Streamable          bool          `json:"streamable"`
	Previewable         bool          `json:"previewable"`
	Sampleable          bool          `json:"sampleable"`
	Downloadable        bool          `json:"downloadable"`
	Displayable         bool          `json:"displayable"`
	PurchasableAt       int           `json:"purchasable_at"`
	StreamableAt        int           `json:"streamable_at"`
	Hires               bool          `json:"hires"`
	HiresStreamable     bool          `json:"hires_streamable"`
	Awards              []interface{} `json:"awards"`
	Goodies             []interface{} `json:"goodies"`
	Area                interface{}   `json:"area"`
	Catchline           string        `json:"catchline"`
	Composer            struct {
		Id          int         `json:"id"`
		Name        string      `json:"name"`
		Slug        string      `json:"slug"`
		AlbumsCount int         `json:"albums_count"`
		Picture     interface{} `json:"picture"`
		Image       interface{} `json:"image"`
	} `json:"composer"`
	CreatedAt                      int           `json:"created_at"`
	GenresList                     []string      `json:"genres_list"`
	Period                         interface{}   `json:"period"`
	Copyright                      string        `json:"copyright"`
	IsOfficial                     bool          `json:"is_official"`
	MaximumTechnicalSpecifications string        `json:"maximum_technical_specifications"`
	ProductSalesFactorsMonthly     int           `json:"product_sales_factors_monthly"`
	ProductSalesFactorsWeekly      int           `json:"product_sales_factors_weekly"`
	ProductSalesFactorsYearly      int           `json:"product_sales_factors_yearly"`
	ProductType                    string        `json:"product_type"`
	ProductUrl                     string        `json:"product_url"`
	RecordingInformation           string        `json:"recording_information"`
	RelativeUrl                    string        `json:"relative_url"`
	ReleaseTags                    []interface{} `json:"release_tags"`
	ReleaseType                    string        `json:"release_type"`
	Slug                           string        `json:"slug"`
	Subtitle                       string        `json:"subtitle"`
	Tracks                         struct {
		Offset int `json:"offset"`
		Limit  int `json:"limit"`
		Total  int `json:"total"`
		Items  []struct {
			MaximumBitDepth int    `json:"maximum_bit_depth"`
			Copyright       string `json:"copyright"`
			Performers      string `json:"performers"`
			AudioInfo       struct {
				ReplaygainTrackPeak float64 `json:"replaygain_track_peak"`
				ReplaygainTrackGain float64 `json:"replaygain_track_gain"`
			} `json:"audio_info"`
			Performer struct {
				Name string `json:"name"`
				Id   int    `json:"id"`
			} `json:"performer"`
			Work     interface{} `json:"work"`
			Composer struct {
				Name string `json:"name"`
				Id   int    `json:"id"`
			} `json:"composer"`
			Isrc                string      `json:"isrc"`
			Title               string      `json:"title"`
			Version             string      `json:"version"`
			Duration            int         `json:"duration"`
			ParentalWarning     bool        `json:"parental_warning"`
			TrackNumber         int         `json:"track_number"`
			MaximumChannelCount int         `json:"maximum_channel_count"`
			Id                  int         `json:"id"`
			MediaNumber         int         `json:"media_number"`
			MaximumSamplingRate float64     `json:"maximum_sampling_rate"`
			ReleaseDateOriginal interface{} `json:"release_date_original"`
			ReleaseDateDownload interface{} `json:"release_date_download"`
			ReleaseDateStream   interface{} `json:"release_date_stream"`
			ReleaseDatePurchase interface{} `json:"release_date_purchase"`
			Purchasable         bool        `json:"purchasable"`
			Streamable          bool        `json:"streamable"`
			Previewable         bool        `json:"previewable"`
			Sampleable          bool        `json:"sampleable"`
			Downloadable        bool        `json:"downloadable"`
			Displayable         bool        `json:"displayable"`
			PurchasableAt       int         `json:"purchasable_at"`
			StreamableAt        int         `json:"streamable_at"`
			Hires               bool        `json:"hires"`
			HiresStreamable     bool        `json:"hires_streamable"`
		} `json:"items"`
	} `json:"tracks"`
	Description string `json:"description"`
}

func (m GetAlbumMetaResponse) getTitle() string {
	title := m.Title
	if m.Version != "" {
		title += " (" + m.Version + ")"
	}
	return title
}

type AlbumAttributes struct {
	Album        string
	Artist       string
	Year         string
	BitDepth     int
	SamplingRate float64
}

func (m GetAlbumMetaResponse) GetAlbumAttributes(fileUrl GetFileURLResponse) AlbumAttributes {
	return AlbumAttributes{
		Album:        m.getTitle(),
		Artist:       m.Artist.Name,
		Year:         m.ReleaseDateOriginal,
		BitDepth:     fileUrl.BitDepth,
		SamplingRate: fileUrl.SamplingRate,
	}
}

func (attributes AlbumAttributes) folderFormat() string {
	year := strings.Split(attributes.Year, "-")[0]
	return fmt.Sprintf("%s - %s (%s) [%dB-%.1fkHz]", attributes.Artist, attributes.Album, year, attributes.BitDepth, attributes.SamplingRate)
}
