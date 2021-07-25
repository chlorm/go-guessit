// Copyright 2017 Cody Opel <codyopel@gmail.com>
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

// TODO: add support for guessit-rest

package guessit

import (
	"encoding/json"
	"os/exec"
)

// http://guessit.readthedocs.io/en/latest/properties.html
type Properties struct {
	// Main properties
	Type []struct {
		Value string `json:"value"`
		Raw   string `json:"raw"`
		Start int    `json:"start"`
		End   int    `json:"end"`
	} `json:"type"`
	Title []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"title,omitempty"`
	EpisodeTitle []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"episode_title,omitempty"`
	AlternativeTitle []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"alternative_title,omitempty"`
	Container []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"container,omitempty"`
	MimeType []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"mimetype,omitempty"`
	Date []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"date,omitempty"`
	Year []struct {
		Value int    `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"year,omitempty"`
	ReleaseGroup []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"release_group,omitempty"`
	Website []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"website,omitempty"`
	StreamingService []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"streaming_service,omitempty"`
	// Episode properties
	Season []struct {
		Value int    `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"season,omitempty"`
	Episode []struct {
		Value int    `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"episode,omitempty"`
	EpisodeCount []struct {
		Value int    `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"episode_count,omitempty"`
	SeasonCount []struct {
		Value int    `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"season_count,omitempty"`
	EpisodeDetails []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"episode_details,omitempty"`
	EpisodeFormat []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"episode_format,omitempty"`
	Part []struct {
		Value int    `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"part,omitempty"`
	Version []struct {
		Value int    `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"version,omitempty"`
	//Video properties
	Format []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"format,omitempty"`
	ScreenSize []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"screen_size,omitempty"`
	VideoCodec []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"video_codec,omitempty"`
	VideoProfile []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"video_profile,omitempty"`
	VideoApi []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"video_api,omitempty"`
	// Audio properties
	AudioChannels []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"audio_channels,omitempty"`
	AudioCodec []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"audio_codec,omitempty"`
	AudioProfile []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"audio_profile,omitempty"`
	// Localization properties
	Country []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"country,omitempty"`
	Language []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"language,omitempty"`
	SubtitleLanguage []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"subtitle_language,omitempty"`
	// Other properties
	Bonus []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"bonus,omitempty"`
	BonusTitle []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"bonus_title,omitempty"`
	Cd []struct {
		Value int    `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"cd,omitempty"`
	CdCount []struct {
		Value int    `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"cd_count,omitempty"`
	Crc32 []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"crc32,omitempty"`
	Uuid []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"uuid,omitempty"`
	Size []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"size,omitempty"`
	Edition []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"edition,omitempty"`
	Film []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"film,omitempty"`
	FilmTitle []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"film_title,omitempty"`
	FilmSeries []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"film_series,omitempty"`
	Other []struct {
		Value string `json:"value,omitempty"`
		Raw   string `json:"raw,omitempty"`
		Start int    `json:"start,omitempty"`
		End   int    `json:"end,omitempty"`
	} `json:"other,omitempty"`
}

// Returns the output of Guessit (advanced) as a struct
func Guessit(filename string) (*Properties, error) {
	guessit, err := exec.LookPath("guessit")
	if err != nil {
		return nil, err
	}

	guessitOut, err := exec.Command(guessit, "--json", "--advanced", "--enforce-list", filename).Output()
	if err != nil {
		return nil, err
	}

	// Unmarshal JSON into struct
	var properties *Properties
	if err = json.Unmarshal(guessitOut, &properties); err != nil {
		return nil, err
	}

	return properties, nil
}
