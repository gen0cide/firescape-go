package models

import "github.com/jinzhu/gorm"

type Advertisement struct {
	gorm.Model
	Message string
}

type ShopDef struct {
	gorm.Model
	Name         string
	General      bool
	SellModifier int
	buyModifier  int
	MinX         int
	MaxX         int
	MinY         int
	MaxY         int
	RespawnRate  int
	Greeting     string
	Options      []string
	Items        []*ShopItem
}

type ShopItem struct {
	gorm.Model
	Amount  int
	ItemDef *ItemDef
	ShopDef *ShopDef
}

type NPCLocation struct {
	gorm.Model
	NPCDef *NPCDef
	StartX int
	StartY int
	MinX   int
	MaxX   int
	MinY   int
	MaxY   int
}

type NPCDef struct {
	gorm.Model
	Name         string
	Description  string
	Command      string
	Attack       int
	Strength     int
	Hits         int
	Defense      int
	Attackable   bool
	Aggressive   bool
	RespawnTime  int
	Sprites      []int
	HairColor    int
	TopColor     int
	BottomColor  int
	SkinColor    int
	Camera1      int
	Camera2      int
	WalkModel    int
	CombatModel  int
	CombatSprite int
}

type NPCDrop struct {
	gorm.Model
	Amount  int
	Weight  int
	ItemDef *ItemDef
}

type ItemDef struct {
	gorm.Model
	Name        string
	Description string
	Command     string
	Sprite      int
	BasePrice   int
	Stackable   bool
	Wieldable   bool
	PictureMask bool
}

type GameObjectDef struct {
	gorm.Model
	Name          string
	Description   string
	Command1      string
	Command2      string
	Type          int
	Width         int
	Height        int
	GroundItemVar int
	ObjectModel   string
}

type PrayerDef struct {
	gorm.Model
	Name          string
	RequiredLevel int
	DrainRate     int
	Description   string
}

type SpellDef struct {
	gorm.Model
	Name            string
	Description     string
	RequiredLevel   int
	Type            int
	RuneCount       int
	RequiredRunes   []*RequiredRuneDef
	Experience      int
	AggressiveLevel int
}

type RequiredRuneDef struct {
	gorm.Model
	Rune     *ItemDef
	Quantity int
}

type TileDef struct {
	gorm.Model
	Color      int
	Unknown    int
	ObjectType int
}

type FireMakingDef struct {
	gorm.Model
	Log        *ItemDef
	Level      int
	Experience int
	Length     int
}

type ArrowMakingDef struct {
	gorm.Model
	ArrowHead     *ItemDef
	Arrow         *ItemDef
	RequiredLevel int
	Experience    int
}

type BowMakingDef struct {
	gorm.Model
	UnstrungBow   *ItemDef
	Bow           *ItemDef
	RequiredLevel int
	Experience    int
}

type CookingDef struct {
	gorm.Model
	RawFood       *ItemDef
	CookedFood    *ItemDef
	BurnedFood    *ItemDef
	RequiredLevel int
	Experience    int
}

type CraftingDef struct {
	gorm.Model
	Item          *ItemDef
	RequiredLevel int
	Experience    int
}

type DartTipDef struct {
	gorm.Model
	DartTip       *ItemDef
	Dart          *ItemDef
	RequiredLevel int
	Experience    int
}

type HealingDef struct {
	gorm.Model
	Item  *ItemDef
	Heals int
}

type GemCuttingDef struct {
	gorm.Model
	UncutGem      *ItemDef
	Gem           *ItemDef
	RequiredLevel int
	Experience    int
}

type PotionDef struct {
	gorm.Model
	Ingredient    *ItemDef
	Potion        *ItemDef
	RequiredLevel int
	Experience    int
}

type AdvancedPotionDef struct {
	gorm.Model
	Ingredient       *ItemDef
	UnfinishedPotion *ItemDef
	Potion           *ItemDef
	RequiredLevel    int
	Experience       int
}

type LogCuttingDef struct {
	gorm.Model
	Log                *ItemDef
	ShaftLevel         int
	ShaftAmount        int
	Shortbow           *ItemDef
	ShortbowLevel      int
	ShortbowExperience int
	Longbow            *ItemDef
	LongbowLevel       int
	LongbowExperience  int
}

type SmeltingDef struct {
	gorm.Model
	Ore           *ItemDef
	Bar           *ItemDef
	RequiredLevel int
	Experience    int
	RequiredOres  []*RequiredOreDef
}

type RequiredOreDef struct {
	gorm.Model
	Ore    *ItemDef
	Amount int
}

type SmithingDef struct {
	gorm.Model
	RequiredLevel int
	BarsRequired  int
	Item          *ItemDef
	Amount        int
}

type HerbDef struct {
	gorm.Model
	HiddenHerb    *ItemDef
	RealHerb      *ItemDef
	RequiredLevel int
	Experience    int
}

type WeieldableDef struct {
	gorm.Model
	Item          *ItemDef
	Sprite        int
	Type          int
	WieldPosition int
	Armor         int
	WeaponAim     int
	WeaponPower   int
	Magic         int
	Prayer        int
	Range         int
	FemaleOnly    bool
	RequiredStats []*RequiredStatDef
}

type RequiredStatDef struct {
	gorm.Model
	Stat          *StatDef
	RequiredLevel int
}

type StatDef struct {
	gorm.Model
	Name string
}

type ChestLootDef struct {
	gorm.Model
	Chest  *GameObjectDef
	Item   *ItemDef
	Amount int
}

type CerterDef struct {
	gorm.Model
	NPC  *NPCDef
	Type string
	Name string
	Cert *ItemDef
	Item *ItemDef
}

type FishingDef struct {
	gorm.Model
	Target        *GameObjectDef
	Fish          *ItemDef
	Net           *ItemDef
	Bait          *ItemDef
	RequiredLevel int
	Experience    int
}

type MiningDef struct {
	gorm.Model
	Rock          *GameObjectDef
	Ore           *ItemDef
	RequiredLevel int
	Experience    int
	RespawnTime   int
}

type WoodcuttingDef struct {
	gorm.Model
	Tree          *GameObjectDef
	Log           *ItemDef
	Fell          int
	RequiredLevel int
	Experience    int
	RespawnTime   int
}

type WallObject struct {
	gorm.Model
	Name        string
	Description string
	Command1    string
	Command2    string
	ModelVar1   int
	ModelVar2   int
	ModelVar3   int
	DoorType    int
	Unknown     int
}

type ObjectTeleportDef struct {
	gorm.Model
	StartX  int
	StartY  int
	Command string
	EndX    int
	EndY    int
}

type GameObjectLocation struct {
	gorm.Model
	Object    *GameObjectDef
	X         int
	Y         int
	Direction int
	Type      int
}

type ItemLocation struct {
	gorm.Model
	Item        *ItemDef
	X           int
	Y           int
	Amount      int
	RespawnTime int
}

type ActiveTile struct {
	gorm.Model
	Players []*Player
	NPCs    []*NPC
}

type Player struct {
	gorm.Model
	Mob
	Username      string
	UsernameHash  string
	Password      string
	LoggedIn      bool
	LastPing      uint64
	Rank          int
	Appearance    *PlayerAppearance
	Stats         []*PlayerStat
	Suspicious    bool
	Inventory     *Inventory
	Bank          *Bank
	DrainRate     int
	InBank        bool
	SessionKeys   *PlayerSession
	CombatStyle   int
	LastCharge    uint64
	LastReport    uint64
	PVPAttacks    []*PVPAttack
	LastSpellCast uint64
	MaleGender    bool
	CurrentIP     string
	LastIP        string
	LoginTime     uint64
	LastLoginTime uint64
	FriendList    []*Player
	IgnoreList    []*Player
	TradeOffer    []*InvItem
	DuelOffer     []*InvItem
}

type PlayerAppearance struct {
	gorm.Model
	HairColor int
	TopColor  int
	PantColor int
	SkinColor int
	Head      int
	Body      int
}

type Inventory struct {
	Player *Player
	Items  []*InvItem
}

type PVPAttack struct {
	gorm.Model
	Attacker  *Player
	Victim    *Player
	Timestamp uint64
}

type PlayerSession struct {
	gorm.Model
	Player *Player
	KeyA   int
	KeyB   int
	KeyC   int
	KeyD   int
}

type Bank struct {
	Items []*InvItem
}

type PlayerStat struct {
	gorm.Model
	Player     *Player
	Stat       *StatDef
	Current    int
	Max        int
	Experience int
}

type PlayerSetting struct {
	gorm.Model
	SettingID int
	Player    *Player
	Value     bool
}

type Entity struct {
	World    *World
	Location *Point
	ID       int
	Index    int
}

type InvItem struct {
	Entity
	Amount  int
	Wielded bool
}

type World struct {
	MaxWidth    int
	MaxHeight   int
	ActiveTiles [][]*ActiveTile
	TileValues  [][]*TileValue
	Jackpot     int
	Players     []*Player
	NPCs        []*NPC
	Shops       []*Shop
}

type Mob struct {
	Entity
	Sprites           [][]int
	MobSprite         int
	HasMoved          bool
	CombatLevel       int
	AppearanceChanged bool
	AppearanceID      int
	LastMovement      uint64
	WarnedToMove      bool
	Busy              bool
	SpriteChanged     bool
	CombatWith        *Mob
	CombatTimer       uint64
	LastDamage        int
	ViewArea          *ViewArea
	ActivatedPrayers  []*PrayerDef
	HitsMade          int
	Removed           bool
	LastCombatState   int
}

type NPC struct {
	World *World
	NPCLocation
	NPCDef
	CurrentHitPoints int
	CurrentAttack    int
	CurrentStrength  int
	CurrentDefense   int
	Blocker          *Player
	ShouldRespawn    bool
}

type ViewArea struct {
	World *World
	Mob   *Mob
}

type TileValue struct {
	MapValue    int
	ObjectValue int
}

type Tile struct {
	GroundElevation int
	GroundTexture   int
	RoofTexture     int
	HorizontalWall  int
	VerticalWall    int
	DiagonalWall    int
	GroundOverlay   int
}

type TelePoint struct {
	Point
	Command string
}

type Shop struct {
	World *World
	ShopDef
}

type Sector struct {
	Tiles []*Tile
}

type Projectile struct {
	Caster *Mob
	Victim *Mob
	Type   int
}

type Point struct {
	X int
	Y int
}

type Item struct {
	World       *World
	Owner       *Player
	Amount      int
	SpawnedTime uint64
	Removed     bool
	Location    *ItemLocation
}

type GameObject struct {
	Entity
	GameObjectLocation
	Direction int
	Type      int
	Grainable bool
	Removed   bool
}

type ChatMessage struct {
	Sender    *Player
	Recipient *Player
	Message   string
}

type Bubble struct {
	Owner *Player
	Item  *ItemDef
}
