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

  "github.com/chlorm/go-sys-utils/findbinpath"
)

// http://guessit.readthedocs.io/en/latest/properties.html
type Properties struct {
  // Main properties
  Type struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"type"`
  Title struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"title"`
  AlternativeTitle struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"alternative_title"`
  Container struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"container"`
  MimeType struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"mimetype"`
  Date struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"date"`
  Year struct {
    Value int `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"year"`
  ReleaseGroup struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"release_group"`
  Website struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"website"`
  StreamingService struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"streaming_service"`
  // Episode properties
  Season struct {
    Value int `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"season"`
  Episode struct {
    Value int `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"episode"`
  EpisodeCount struct {
    Value int `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"episode_count"`
  SeasonCount struct {
    Value int `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"season_count"`
  EpisodeDetails struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"episode_details"`
  EpisodeFormat struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"episode_format"`
  Part struct {
    Value int `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"part"`
  Version struct {
    Value int `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"version"`
  //Video properties
  Format struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"format"`
  ScreenSize struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"screen_size"`
  VideoCodec struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"video_codec"`
  VideoProfile struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"video_profile"`
  VideoApi struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"video_api"`
  // Audio properties
  AudioChannels struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"audio_channels"`
  AudioCodec struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"audio_codec"`
  AudioProfile struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"audio_profile"`
  // Localization properties
  Country struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"country"`
  Language struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"language"`
  SubtitleLanguage struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"subtitle_language"`
  // Other properties
  Bonus struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"bonus"`
  BonusTitle struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"bonus_title"`
  Cd struct {
    Value int `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"cd"`
  CdCount struct {
    Value int `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"cd_count"`
  Crc32 struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"crc32"`
  Uuid struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"uuid"`
  Size struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"size"`
  Edition struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"edition"`
  Film struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"film"`
  FilmTitle struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"film_title"`
  FilmSeries struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"film_series"`
  Other struct {
    Value string `json:"value"`
    Raw string `json:"raw"`
    Start int `json:"start"`
    End int `json:"end"`
  } `json:"other"`
}

// Call guessit CLI
func Guessit(filename string) (*Properties, error) {
  var err error

  _, err = findbinpath.FindBinPath("guessit")
  if err != nil {
    return nil, err
  }

  guessitOut, err := exec.Command("guessit", "-j", "-a", filename).Output()
  if err != nil {
    return nil, err
  }

  // Unmarshal JSON into struct
  var properties *Properties
  err = json.Unmarshal(guessitOut, &properties)
  if err != nil {
    return nil, err
  }

  return properties, nil
}
