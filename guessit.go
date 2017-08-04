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

// TODO:
// - add support for guessit-rest
// - allow passing CLI additional arguments

package guessit

import (
  "encoding/json"
  "os/exec"

  "github.com/chlorm/go-sys-utils"
)

// http://guessit.readthedocs.io/en/latest/properties.html
type Properties struct {
  // Main properties
  Type string `json:"type"`
  Title string `json:"title,omitempty"`
  AlternativeTitle string `json:"alternative_title,omitempty"`
  Container string `json:"container,omitempty"`
  MimeType string `json:"mimetype,omitempty"`
  Date string `json:"date,omitempty"`
  Year int `json:"year,omitempty"`
  ReleaseGroup string `json:"release_group,omitempty"`
  Website string `json:"website,omitempty"`
  StreamingService string `json:"streaming_service,omitempty"`
  // Episode properties
  Season int `json:"season,omitempty"`
  Episode int `json:"episode,omitempty"`
  EpisodeCount int `json:"episode_count,omitempty"`
  SeasonCount int `json:"season_count,omitempty"`
  EpisodeDetails string `json:"episode_details,omitempty"`
  EpisodeFormat string `json:"episode_format,omitempty"`
  Part int `json:"part,omitempty"`
  Version int `json:"version,omitempty"`
  //Video properties
  Format string `json:"format,omitempty"`
  ScreenSize string `json:"screen_size,omitempty"`
  VideoCodec string `json:"video_codec,omitempty"`
  VideoProfile string `json:"video_profile,omitempty"`
  VideoApi string `json:"video_api,omitempty"`
  // Audio properties
  AudioChannels string `json:"audio_channels,omitempty"`
  AudioCodec string `json:"audio_codec,omitempty"`
  AudioProfile string `json:"audio_profile,omitempty"`
  // Localization properties
  Country string `json:"country,omitempty"`
  Language string `json:"language,omitempty"`
  SubtitleLanguage string `json:"subtitle_language,omitempty"`
  // Other properties
  Bonus string `json:"bonus,omitempty"`
  BonusTitle string `json:"bonus_title,omitempty"`
  Cd int `json:"cd,omitempty"`
  CdCount int `json:"cd_count,omitempty"`
  Crc32 string `json:"crc32,omitempty"`
  Uuid string `json:"uuid,omitempty"`
  Size string `json:"size,omitempty"`
  Edition string `json:"edition,omitempty"`
  Film string `json:"film,omitempty"`
  FilmTitle string `json:"film_title,omitempty"`
  FilmSeries string `json:"film_series,omitempty"`
  Other string `json:"other,omitempty"`
}

// Returns the output of Guessit as a struct
func Guessit(filename string) (*Properties, error) {
  var err error

  if _, err = sysutils.SearchEnvPath("PATH", "guessit", ":"); err != nil {
    return nil, err
  }

  guessitOut, err := exec.Command("guessit", "-j", filename).Output()
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
