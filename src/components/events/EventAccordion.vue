<template>
    <div>
        <b-card no-body class="mb-1" v-for="(category, categoryId) in categories" :key="categoryId">
            <b-card-header header-tag="header" class="p-1" role="tab">
                <b-btn block href="#" v-b-toggle="accordionName + '-' + categoryId" variant="secondary">
                    {{category.name}}
                </b-btn>
            </b-card-header>
            <b-collapse :id="accordionName + '-' + categoryId" :accordion="accordionName" role="tabpanel">
                <b-card-body>
                    <component :is="category.component"/>
                </b-card-body>
            </b-collapse>
        </b-card>
    </div>
</template>

<script>
    import AttackerDoubleTouchedBall from "./inputs/AttackerDoubleTouchedBall"
    import AttackerInDefenseArea from "./inputs/AttackerInDefenseArea"
    import AttackerTooCloseToDefenseArea from "./inputs/AttackerTooCloseToDefenseArea"
    import AttackerTouchedKeeper from "./inputs/AttackerTouchedKeeper"
    import BallLeftFieldGoalLine from "./inputs/BallLeftFieldGoalLine"
    import BallLeftFieldTouchLine from "./inputs/BallLeftFieldTouchLine"
    import BotCrashDrawn from "./inputs/BotCrashDrawn"
    import BotCrashUnique from "./inputs/BotCrashUnique"
    import BotCrashUniqueSkipped from "./inputs/BotCrashUniqueSkipped"
    import BotDribbledBallTooFar from "./inputs/BotDribbledBallTooFar"
    import BotHeldBallDeliberately from "./inputs/BotHeldBallDeliberately"
    import BotInterferedPlacement from "./inputs/BotInterferedPlacement"
    import BotKickedBallTooFast from "./inputs/BotKickedBallTooFast"
    import BotPushedBot from "./inputs/BotPushedBot"
    import BotPushedBotSkipped from "./inputs/BotPushedBotSkipped"
    import BotTippedOver from "./inputs/BotTippedOver"
    import BotTooFastInStop from "./inputs/BotTooFastInStop"
    import ChippedGoal from "./inputs/ChippedGoal"
    import DefenderInDefenseArea from "./inputs/DefenderInDefenseArea"
    import DefenderInDefenseAreaPartially from "./inputs/DefenderInDefenseAreaPartially"
    import DefenderTooCloseToKickPoint from "./inputs/DefenderTooCloseToKickPoint"
    import Goal from "./inputs/Goal"
    import AimlessKick from "./inputs/AimlessKick"
    import IndirectGoal from "./inputs/IndirectGoal"
    import KeeperHeldBall from "./inputs/KeeperHeldBall"
    import KickTimeout from "./inputs/KickTimeout"
    import MultipleFouls from "./inputs/MultipleFouls"
    import MultipleCards from "./inputs/MultipleCards"
    import MultiplePlacementFailures from "./inputs/MultiplePlacementFailures"
    import NoProgressInGame from "./inputs/NoProgressInGame"
    import PlacementFailedByOpponent from "./inputs/PlacementFailedByOpponent"
    import PlacementFailedByTeamInFavor from "./inputs/PlacementFailedByTeamInFavor"
    import UnsportingBehaviorMajor from "./inputs/UnsportingBehaviorMajor"
    import UnsportingBehaviorMinor from "./inputs/UnsportingBehaviorMinor"
    import PlacementSucceeded from "./inputs/PlacementSucceeded"
    import Prepared from "./inputs/Prepared"

    export default {
        name: "EventAccordion",
        components: {
            AttackerDoubleTouchedBall,
            AttackerInDefenseArea,
            AttackerTooCloseToDefenseArea,
            AttackerTouchedKeeper,
            BallLeftFieldGoalLine,
            BallLeftFieldTouchLine,
            BotCrashDrawn,
            BotCrashUnique,
            BotCrashUniqueSkipped,
            BotDribbledBallTooFar,
            BotHeldBallDeliberately,
            BotInterferedPlacement,
            BotKickedBallTooFast,
            BotPushedBot,
            BotPushedBotSkipped,
            BotTippedOver,
            BotTooFastInStop,
            ChippedGoal,
            DefenderInDefenseArea,
            DefenderInDefenseAreaPartially,
            DefenderTooCloseToKickPoint,
            Goal,
            AimlessKick,
            IndirectGoal,
            KeeperHeldBall,
            KickTimeout,
            MultipleFouls,
            MultipleCards,
            MultiplePlacementFailures,
            NoProgressInGame,
            PlacementFailedByOpponent,
            PlacementFailedByTeamInFavor,
            UnsportingBehaviorMajor,
            UnsportingBehaviorMinor,
            PlacementSucceeded,
            Prepared,
        },
        props: {
            accordionName: {
                type: String
            },
            categories: {
                type: Array
            }
        }
    }
</script>

<style scoped>

</style>