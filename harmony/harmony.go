package harmony

import (
	"bytes"
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/mritd/logger"
)

const TokenPrefix = "harmony:"

const (
	harmonyProjectID  = "101653523863572770"
	harmonyKeyID      = "1a2c9713eec4432e89527627edb93e98"
	harmonySubAccount = "116892607"
	harmonyTokenURI   = "https://oauth-login.cloud.huawei.com/oauth2/v3/token"
	harmonyPrivateKey = `-----BEGIN PRIVATE KEY-----
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
-----END PRIVATE KEY-----`
)

type PushMessage struct {
	DeviceToken string
	Title       string
	Body        string
	ExtParams   map[string]interface{}
	IsDelete    bool
}

type serviceAccountKey struct {
	ProjectID  string `json:"project_id"`
	KeyID      string `json:"key_id"`
	PrivateKey string `json:"private_key"`
	SubAccount string `json:"sub_account"`
	TokenURI   string `json:"token_uri"`
}

type jwtSupplier struct {
	mu        sync.Mutex
	token     string
	expiresAt time.Time
	sa        serviceAccountKey
	key       *rsa.PrivateKey
	duration  time.Duration
}

type client struct {
	httpClient *http.Client
	messageURL string
	jwt        *jwtSupplier
}

type clickAction struct {
	ActionType int `json:"actionType"`
}

type badge struct {
	AddNum int `json:"addNum,omitempty"`
	SetNum int `json:"setNum,omitempty"`
}

type notification struct {
	Category       string      `json:"category,omitempty"`
	Title          string      `json:"title,omitempty"`
	Body           string      `json:"body,omitempty"`
	ClickAction    clickAction `json:"clickAction"`
	Style          int         `json:"style,omitempty"`
	InboxContent   []string    `json:"inboxContent,omitempty"`
	Badge          *badge      `json:"badge,omitempty"`
	Image          string      `json:"image,omitempty"`
	Sound          string      `json:"sound,omitempty"`
	ForegroundShow *bool       `json:"foregroundShow,omitempty"`
	ExtraData      interface{} `json:"extraData,omitempty"`
}

type alertPayload struct {
	Notification notification `json:"notification"`
}

type backgroundPayload struct {
	ExtraData string `json:"extraData"`
}

type pushOptions struct {
	Ttl         int64 `json:"ttl,omitempty"`
	TestMessage bool  `json:"testMessage,omitempty"`
}

type target struct {
	Token []string `json:"token"`
}

type message struct {
	Payload     interface{}  `json:"payload"`
	Target      target       `json:"target"`
	PushOptions *pushOptions `json:"pushOptions,omitempty"`
}

var (
	clients       = make(chan *client, 1)
	initialized   bool
	messageURL    string
	tokenSupplier *jwtSupplier
)

func Init(maxClientCount int, pushServerDomain string) error {
	if maxClientCount < 1 {
		return fmt.Errorf("invalid number of clients")
	}
	sa, err := loadServiceAccountKey()
	if err != nil {
		return err
	}
	if pushServerDomain == "" {
		pushServerDomain = "push-api.cloud.huawei.com"
	}
	messageURL = fmt.Sprintf("https://%s/v3/%s/messages:send", pushServerDomain, sa.ProjectID)
	tokenSupplier, err = newJWTSupplier(*sa, time.Hour)
	if err != nil {
		return err
	}
	clients = make(chan *client, maxClientCount)
	for i := 0; i < min(runtime.NumCPU(), maxClientCount); i++ {
		clients <- &client{
			httpClient: &http.Client{
				Timeout: 5 * time.Second,
			},
			messageURL: messageURL,
			jwt:        tokenSupplier,
		}
	}
	initialized = true
	logger.Infof("init harmony client success, count=%d", min(runtime.NumCPU(), maxClientCount))
	return nil
}

func IsConfigured() bool {
	return initialized && tokenSupplier != nil && messageURL != ""
}

func Push(msg *PushMessage) (int, error) {
	if !IsConfigured() {
		return 500, errors.New("harmony push is not configured")
	}
	if msg.DeviceToken == "" {
		return 400, errors.New("device token is empty")
	}
	if msg.ExtParams == nil {
		msg.ExtParams = map[string]interface{}{}
	}

	pushType := 0
	var payload interface{}
	if msg.IsDelete {
		pushType = 6
		payload = backgroundPayload{ExtraData: buildExtraData(msg.ExtParams)}
	} else {
		noti := notification{
			Category:    getString(msg.ExtParams, "category", "MARKETING"),
			Title:       msg.Title,
			Body:        msg.Body,
			ClickAction: clickAction{ActionType: 0},
		}
		if image, ok := msg.ExtParams["image"]; ok {
			noti.Image = fmt.Sprint(image)
		}
		if badgeValue := buildBadge(msg.ExtParams); badgeValue != nil {
			noti.Badge = badgeValue
		}
		if sound, ok := msg.ExtParams["sound"]; ok {
			soundValue := strings.TrimSpace(fmt.Sprint(sound))
			if soundValue != "" {
				noti.Sound = soundValue
			}
		}
		if style, ok := toInt(msg.ExtParams["style"]); ok {
			noti.Style = style
		}
		inboxContent := getInboxContent(msg.ExtParams)
		if len(inboxContent) > 0 {
			noti.InboxContent = inboxContent
			if noti.Style == 0 {
				noti.Style = 3
			}
		}
		if fg, ok := toBool(msg.ExtParams["foreground_show"]); ok {
			noti.ForegroundShow = &fg
		}
		payload = alertPayload{Notification: noti}
	}

	options := &pushOptions{TestMessage: true}
	if ttl, ok := toInt64(msg.ExtParams["ttl"]); ok && ttl > 0 {
		options.Ttl = ttl
	}

	reqBody, err := json.Marshal(message{
		Payload:     payload,
		Target:      target{Token: []string{msg.DeviceToken}},
		PushOptions: options,
	})
	if err != nil {
		return 500, err
	}
	token, err := tokenSupplier.Get()
	if err != nil {
		return 500, err
	}
	req, err := http.NewRequest("POST", messageURL, bytes.NewBuffer(reqBody))
	if err != nil {
		return 500, err
	}
	req.Header.Set("Content-Type", "application/json; charset=UTF-8")
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("push-type", strconv.Itoa(pushType))

	client := <-clients
	clients <- client
	resp, err := client.httpClient.Do(req)
	if err != nil {
		return 500, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		return resp.StatusCode, fmt.Errorf(strings.TrimSpace(string(body)))
	}
	return 200, nil
}

func (s *jwtSupplier) Get() (string, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.token != "" && time.Now().Before(s.expiresAt.Add(-30*time.Second)) {
		return s.token, nil
	}
	now := time.Now().UTC()
	iat := now.Unix()
	exp := iat + int64(s.duration.Seconds())
	aud := s.sa.TokenURI
	if aud == "" {
		aud = "https://oauth-login.cloud.huawei.com/oauth2/v3/token"
	}
	claims := jwt.MapClaims{
		"aud": aud,
		"iss": s.sa.SubAccount,
		"exp": exp,
		"iat": iat,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodPS256, claims)
	token.Header["kid"] = s.sa.KeyID
	token.Header["typ"] = "JWT"
	token.Header["alg"] = "PS256"
	signed, err := token.SignedString(s.key)
	if err != nil {
		return "", err
	}
	s.token = signed
	s.expiresAt = time.Unix(exp, 0)
	return signed, nil
}

func newJWTSupplier(sa serviceAccountKey, duration time.Duration) (*jwtSupplier, error) {
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM([]byte(sa.PrivateKey))
	if err != nil {
		return nil, fmt.Errorf("failed to parse private key: %w", err)
	}
	return &jwtSupplier{
		sa:       sa,
		key:      privateKey,
		duration: duration,
	}, nil
}

func loadServiceAccountKey() (*serviceAccountKey, error) {
	sa := serviceAccountKey{
		ProjectID:  harmonyProjectID,
		KeyID:      harmonyKeyID,
		PrivateKey: harmonyPrivateKey,
		SubAccount: harmonySubAccount,
		TokenURI:   harmonyTokenURI,
	}
	if sa.ProjectID == "" || sa.KeyID == "" || sa.SubAccount == "" || sa.PrivateKey == "" {
		return nil, errors.New("invalid service account key file")
	}
	return &sa, nil
}

func buildExtraData(params map[string]interface{}) string {
	if len(params) == 0 {
		return ""
	}
	data, err := json.Marshal(params)
	if err != nil {
		return ""
	}
	return string(data)
}

func getString(params map[string]interface{}, key string, def string) string {
	if val, ok := params[key]; ok {
		str := strings.TrimSpace(fmt.Sprint(val))
		if str != "" {
			return str
		}
	}
	return def
}

func getInboxContent(params map[string]interface{}) []string {
	keys := []string{"inboxContent", "inbox_content", "inboxcontent"}
	for _, key := range keys {
		val, ok := getParam(params, key)
		if !ok {
			continue
		}
		switch v := val.(type) {
		case []string:
			return sanitizeStringSlice(v)
		case []interface{}:
			result := make([]string, 0, len(v))
			for _, item := range v {
				text := strings.TrimSpace(fmt.Sprint(item))
				if text != "" {
					result = append(result, text)
				}
			}
			return result
		case string:
			return splitInboxContent(v)
		default:
			text := strings.TrimSpace(fmt.Sprint(v))
			if text != "" {
				return []string{text}
			}
		}
	}
	return nil
}

func getParam(params map[string]interface{}, key string) (interface{}, bool) {
	for k, v := range params {
		if strings.EqualFold(k, key) {
			return v, true
		}
	}
	return nil, false
}

func sanitizeStringSlice(values []string) []string {
	result := make([]string, 0, len(values))
	for _, item := range values {
		text := strings.TrimSpace(item)
		if text != "" {
			result = append(result, text)
		}
	}
	return result
}

func splitInboxContent(value string) []string {
	raw := strings.TrimSpace(value)
	if raw == "" {
		return nil
	}
	var parts []string
	if strings.Contains(raw, "|") {
		parts = strings.Split(raw, "|")
	} else if strings.Contains(raw, "\n") {
		parts = strings.Split(raw, "\n")
	} else if strings.Contains(raw, ",") {
		parts = strings.Split(raw, ",")
	} else {
		parts = []string{raw}
	}
	return sanitizeStringSlice(parts)
}

func buildBadge(params map[string]interface{}) *badge {
	badgeValue, ok := getParam(params, "badge")
	if ok {
		switch v := badgeValue.(type) {
		case map[string]interface{}:
			addNum, _ := toInt(v["addNum"])
			setNum, _ := toInt(v["setNum"])
			if addNum != 0 || setNum != 0 {
				return &badge{AddNum: addNum, SetNum: setNum}
			}
		default:
			value, ok := toInt(v)
			if ok {
				if value > 0 {
					return &badge{SetNum: value}
				}
				if value == 0 {
					return &badge{SetNum: 0}
				}
			}
		}
	}
	addNum := getIntParam(params, []string{"badge_add", "badgeadd", "addNum", "badgeAdd"})
	setNum := getIntParam(params, []string{"badge_set", "badgeset", "setNum", "badgeSet"})
	if addNum != 0 || setNum != 0 {
		return &badge{AddNum: addNum, SetNum: setNum}
	}
	return nil
}

func getIntParam(params map[string]interface{}, keys []string) int {
	for _, key := range keys {
		val, ok := getParam(params, key)
		if !ok {
			continue
		}
		if parsed, ok := toInt(val); ok {
			return parsed
		}
	}
	return 0
}

func toBool(val interface{}) (bool, bool) {
	switch v := val.(type) {
	case bool:
		return v, true
	case string:
		parsed, err := strconv.ParseBool(v)
		return parsed, err == nil
	default:
		return false, false
	}
}

func toInt(val interface{}) (int, bool) {
	switch v := val.(type) {
	case int:
		return v, true
	case int64:
		return int(v), true
	case float64:
		return int(v), true
	case string:
		parsed, err := strconv.Atoi(v)
		return parsed, err == nil
	default:
		return 0, false
	}
}

func toInt64(val interface{}) (int64, bool) {
	switch v := val.(type) {
	case int:
		return int64(v), true
	case int64:
		return v, true
	case float64:
		return int64(v), true
	case string:
		parsed, err := strconv.ParseInt(v, 10, 64)
		return parsed, err == nil
	default:
		return 0, false
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
